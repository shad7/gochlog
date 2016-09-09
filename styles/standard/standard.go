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
func ConfigHook(raw interface{}) (types.Styler, error) {
	return NewStandardStyler(raw)
}

// NewStandardStyler creates a new Git-based ChangeLog of type Standard
func NewStandardStyler(config interface{}) (*Standard, error) {
	return new(Standard), nil
}

// GetParserOptions as defined by the Standard convention
func (s *Standard) GetParserOptions() types.ParserOptions {
	return types.ParserOptions{}
}

// GetFormatterOptions as defined by the Standard convention
func (s *Standard) GetFormatterOptions() types.FormatterOptions {
	return types.FormatterOptions{
		MainTemplate:  changelogTemplate(),
		HeaderPartial: headerTemplate(),
		CommitPartial: commitTemplate(),
		FooterPartial: footerTemplate(),
	}
}

// Parse is the Standard implmentation of the interface method
// func (s *Standard) Parse(commit types.Commit, opts types.ParserOptions) (types.Commit, error) {
// 	return commit, nil
// }

// Format is the Standard implementation of the interface method
// func (s *Standard) Format(commit types.Commit, context map[string]string, opts types.FormatterOptions) (string, error) {
// 	return "", nil
// }
