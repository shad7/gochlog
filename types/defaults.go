package types

import "github.com/imdario/mergo"

// DefaultParserOptions returns default options for parsing git commits
func DefaultParserOptions() *ParserOptions {
	return &ParserOptions{}
}

// DefaultFormatterOptions returns default options for formatting parsed git commits
func DefaultFormatterOptions() *FormatterOptions {
	return &FormatterOptions{
		Transform: func(c Commit) Commit {
			return c
		},
		GenerateOn: func(c Commit) bool {
			return true
		},
		FinalizeContext: func(context map[string]string) map[string]string {
			return context
		},
	}
}

// ApplyDefaultParserOptions applies default values for any options not set by the Styler
func ApplyDefaultParserOptions(styleOpts *ParserOptions) (*ParserOptions, error) {
	defaults := DefaultParserOptions()
	if err := mergo.Merge(styleOpts, defaults); err != nil {
		return nil, err
	}

	return styleOpts, nil
}

// MergeFormatterOptions applies default values for any options not set by the Styler
func MergeFormatterOptions(styleOpts *FormatterOptions) (*FormatterOptions, error) {
	defaults := DefaultFormatterOptions()
	if err := mergo.Merge(styleOpts, defaults); err != nil {
		return nil, err
	}

	return styleOpts, nil
}
