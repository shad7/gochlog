package types

// Commit contains a parsed git commit
type Commit struct {
	ShortHash string
	Hash      string
	Author    string
	Email     string
	Date      string
	Subject   string
	Message   string
	Tags      []string
}

// GitOptions control the scope of git commits
type GitOptions struct {
	TagPattern string
	Format     string
	From       string
	To         string
}

// ParserOptions controls the parsing of git commits
type ParserOptions struct {
	SubjectPattern   string
	SubjectParts     []string
	ReferenceActions []string
	IssuePrefixes    []string
	NoteKeywords     []string
	FieldPattern     string
	RevertPattern    string
	RevertParts      []string
	MergePattern     string
	MergeParts       []string
}

// FormatterOptions controls the formatting of git commits
type FormatterOptions struct {
	GroupBy         string
	CommitGroupSort string
	CommitsSort     string
	NoteGroupSort   string
	NotesSort       string
	MainTemplate    string
	HeaderPartial   string
	CommitPartial   string
	FooterPartial   string
	Transform       func(Commit) Commit
	GenerateOn      func(Commit) bool
	FinalizeContext func(map[string]string) map[string]string
}
