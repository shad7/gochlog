package standard

//go:generate templify -p standard changelog.tpl
//go:generate templify -p standard header.tpl
//go:generate templify -p standard commit.tpl
//go:generate templify -p standard footer.tpl

import (
	"github.com/shad7/gochlog/styles"
	"github.com/shad7/gochlog/types"
)

func init() {
	styles.RegisterStyler("standard", ConfigHook)
}

// Standard is a styler that implements the Styler interface
type Standard int

// ConfigHook is the hook to register with the Standard style
func ConfigHook(raw interface{}) (styles.Styler, error) {
	return NewStandardStyler(raw)
}

// NewStandardStyler creates a new Git-based ChangeLog of type Standard
func NewStandardStyler(config interface{}) (*Standard, error) {
	return new(Standard), nil
}

// GetParserOptions as defined by the Standard convention
func (s *Standard) GetParserOptions() *types.ParserOptions {
	return &types.ParserOptions{
		SubjectPattern: `^(\w*)(?:\(([\w\$\.\-\* ]*)\))?\: (.*)$`,
		SubjectParts: []string{
			"type",
			"scope",
			"subject",
		},
		ReferenceActions: []string{
			"close",
			"closes",
			"closed",
			"fix",
			"fixes",
			"fixed",
			"resolve",
			"resolves",
			"resolved",
		},
		IssuePrefixes: []string{"#"},
		NoteKeywords:  []string{"BREAKING CHANGE"},
		FieldPattern:  `^-(.*?)-$`,
		RevertPattern: `^Revert\s"([\s\S]*)"\s*This reverts commit (\w*)\.`,
		RevertParts: []string{
			"header",
			"hash",
		},
		MergePattern: "",
		MergeParts:   []string{},
	}
}

// GetFormatterOptions as defined by the Standard convention
func (s *Standard) GetFormatterOptions() *types.FormatterOptions {
	return &types.FormatterOptions{
		MainTemplate:  changelogTemplate(),
		HeaderPartial: headerTemplate(),
		CommitPartial: commitTemplate(),
		FooterPartial: footerTemplate(),
	}
}
