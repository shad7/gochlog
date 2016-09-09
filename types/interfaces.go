package types

// Styler abstracts the parsing git logs and formatting of a Change Log
type Styler interface {
	GetParserOptions() ParserOptions
	GetFormatterOptions() FormatterOptions
	// Parse(commit Commit, opts ParserOptions) (Commit, error)
	// Format(commit Commit, context map[string]string, opts FormatterOptions) (string, error)
}
