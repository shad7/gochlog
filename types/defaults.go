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
func ApplyDefaultParserOptions(opts *ParserOptions, defaults *ParserOptions) (*ParserOptions, error) {
	popts := DefaultParserOptions()
	if err := mergo.Merge(popts, opts); err != nil {
		return nil, err
	}

	if err := mergo.Merge(popts, defaults); err != nil {
		return nil, err
	}

	return popts, nil
}

// ApplyDefaultFormatterOptions applies default values for any options not set by the Styler
func ApplyDefaultFormatterOptions(opts *FormatterOptions, defaults *FormatterOptions) (*FormatterOptions, error) {
	fopts := DefaultFormatterOptions()
	if err := mergo.Merge(fopts, opts); err != nil {
		return nil, err
	}

	if err := mergo.Merge(fopts, defaults); err != nil {
		return nil, err
	}

	return fopts, nil
}
