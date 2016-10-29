package core

import (
	"fmt"
	"regexp"

	"github.com/davecgh/go-spew/spew"

	"github.com/shad7/gochlog/types"
)

func parse(opts *types.ParserOptions, commits []types.Commit) ([]map[string]string, error) {
	data := []map[string]string{}

	subjexp := regexp.MustCompile(opts.SubjectPattern)
	// fieldexp := regexp.MustCompile(opts.FieldPattern)
	// revertexp := regexp.MustCompile(opts.RevertPattern)
	// mergeexp := regexp.MustCompile(opts.MergePattern)

	for i := range commits {
		s := subjexp.FindStringSubmatch(commits[i].Subject)
		if len(s) != 0 {
			fmt.Println(spew.Sdump(s))
		}
	}

	return data, nil
}
