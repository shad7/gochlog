package core

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/shad7/gochlog/types"
)

func filterTags(tagRegexp string, rawTags string) []string {
	validTag := regexp.MustCompile(tagRegexp)
	return validTag.FindAllString(rawTags, -1)
}

// fetchChanges retrieves the raw commit messages as a list of Commit
func fetchChanges(opts *types.GitOptions) ([]types.Commit, error) {

	const numField = 8
	format := "--format=format:%h%x00%H%x00%an%x00%ae%x00%ad%x00%s%x00%b%x00%D%x00"
	dtformat := "--date=short"

	var out []byte
	out, _ = exec.Command("git", "log", format, dtformat).Output()

	fields := strings.Split(strings.TrimSpace(string(out)), "\x00")
	for i, field := range fields {
		fields[i] = strings.TrimLeft(field, "\r\n")
	}

	var commits []types.Commit
	for i := 0; i+numField <= len(fields); i += numField {
		commits = append(commits, types.Commit{
			ShortHash: fields[i],
			Hash:      fields[i+1],
			Author:    fields[i+2],
			Email:     fields[i+3],
			Date:      fields[i+4],
			Subject:   fields[i+5],
			Message:   fields[i+6],
			Tags:      strings.Split(fields[i+7], ","),
		})
	}

	// make sure only the requested tags are included
	for i := range commits {
		commits[i].Tags = filterTags(opts.TagPattern, commits[i].Tags[0])
	}

	return commits, nil
}
