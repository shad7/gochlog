package types

// Commit contains a parsed git commit
type Commit struct {
	ShortHash string   `json:"shorthash"`
	Hash      string   `json:"hash"`
	Author    string   `json:"author"`
	Email     string   `json:"email"`
	Date      string   `json:"date"`
	Subject   string   `json:"subject"`
	Message   string   `json:"body"`
	Tags      []string `json:"tags"`
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
	SubjectMap       map[string]string
	ReferenceActions []string
	IssuePrefixes    []string
	NoteKeywords     []string
	FieldPattern     string
	RevertPattern    string
	RevertMap        map[string]string
	MergePattern     string
	MergeMap         map[string]string
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
