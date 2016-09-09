package core

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"

	gitci "github.com/shad7/gochlog/types"
)

func filterTags(tagRegexp string, rawTags string) []string {
	validTag := regexp.MustCompile(tagRegexp)
	return validTag.FindAllString(rawTags, -1)
}

// fetchChanges retrieves the raw commit messages as a list of Commit
func fetchChanges(opts *gitci.GitOptions) ([]gitci.Commit, error) {

	format := "--format={\"shorthash\":\"%h\", \"hash\":\"%H\", \"author\":\"%an\", \"email\":\"%ae\", \"tags\":[\"%d\"], \"date\":\"%ad\", \"subject\":\"%s\", \"body\":\"%b\"},"
	dtformat := "--date=short"

	var out []byte
	out, _ = exec.Command("git", "log", format, dtformat).Output()

	// convert from individual lines of json to a list of json
	re := regexp.MustCompile(`,\n$`)
	outs := fmt.Sprintf("[%v]", re.ReplaceAllString(string(out), ""))

	var commits []gitci.Commit
	if err := json.Unmarshal([]byte(outs), &commits); err != nil {
		return nil, err
	}

	// make sure only the requested tags are included
	for i := range commits {
		commits[i].Tags = filterTags(opts.TagPattern, commits[i].Tags[0])
	}

	return commits, nil
}
