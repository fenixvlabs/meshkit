package errorutil

type GlobalFlags struct {
	Verbose                  bool
	RootDir, OutDir, InfoDir string
	SkipDirs                 []string
}

type Info struct {
	Name          string `yaml:"name" json:"name"`                       // the name of the component, e.g. "kuma"
	Type          string `yaml:"type" json:"type"`                       // the type of the component, e.g. "adapter"
	NextErrorCode int    `yaml:"next_error_code" json:"next_error_code"` // the next error code to use. this value will be updated automatically.
	file          string // the path of the component_info.json file
	OldCode       string `yaml:"old_code" json:"old_code"`
	Code          string `yaml:"code" json:"code"`
	CodeIsLiteral bool   `yaml:"code_is_literal" json:"code_is_literal"`
	CodeIsInt     bool   `yaml:"code_is_int" json:"code_is_int"`
	Path          string `yaml:"path" json:"path"`
}

type InfoAll struct {
	Entries              []Info             `yaml:"entries" json:"entries"`                                // raw entries
	LiteralCodes         map[string][]Info  `yaml:"literal_codes" json:"literal_codes"`                    // entries with literal codes
	CallExprCodes        []Info             `yaml:"call_expr_codes" json:"call_expr_codes"`                // entries with call expressions
	DeprecatedNewDefault []string           `yaml:"deprecated_new_default" json:"deprecated_new_default" ` // list of files with usage of deprecated NewDefault func
	Errors               map[string][]Error `yaml:"errors_raw" json:"errors_raw"`                          // map of detected errors created using errors.New(...). The key is the error name, more than 1 entry in the list is a duplication error.
}

func NewInfoAll() *InfoAll {
	return &InfoAll{
		Entries:              []Info{},
		LiteralCodes:         make(map[string][]Info),
		CallExprCodes:        []Info{},
		DeprecatedNewDefault: []string{},
		Errors:               map[string][]Error{}}
}

// Error is used to export Error for e.g. documentation purposes.
// Type Error (errors/types.go) is not reused in order to avoid tight coupling between code and documentation of errors, e.g. on Meshery website.
// It is good practice not to use internal data types in integrations; one should in general transform between internal and external models.
// DDD calls this anti-corruption layer.
// One reason is that one might like to have a different representation externally, e.g. severity 'info' instead of '1'.
// Another one is that it is often desirable to be able to change the internal representation without the need for the consumer
// (in this case, the meshery doc) to have to adjust quickly in order to be able to handle updated content.
// The lifecycles of producers and consumers should not be tightly coupled.
type Error struct {
	Name                 string `yaml:"name" json:"name"`                                   // the name of the error code variable, e.g. "ErrInstallMesh", not guaranteed to be unique as it is package scoped
	Code                 string `yaml:"code" json:"code"`                                   // the code, an int, but exported as string, e.g. "1001", guaranteed to be unique per component-type:component-name
	Severity             string `yaml:"severity" json:"severity"`                           // a textual representation of the type Severity (errors/types.go), i.e. "none", "alert", etc
	LongDescription      string `yaml:"long_description" json:"long_description"`           // might contain newlines (JSON encoded)
	ShortDescription     string `yaml:"short_description" json:"short_description"`         // might contain newlines (JSON encoded)
	ProbableCause        string `yaml:"probable_cause" json:"probable_cause"`               // might contain newlines (JSON encoded)
	SuggestedRemediation string `yaml:"suggested_remediation" json:"suggested_remediation"` // might contain newlines (JSON encoded)
}

// ExternalAll is used to export all Errors including information about the component for e.g. documentation purposes.
type ExternalAll struct {
	ComponentName string           `yaml:"component_name" json:"component_name"` // component type, e.g. "adapter"
	ComponentType string           `yaml:"component_type" json:"component_type"` // component name, e.g. "kuma"
	Errors        map[string]Error `yaml:"errors" json:"errors"`                 // map of all errors with key = code
}

type AnalysisSummary struct {
	MinCode              int                 `yaml:"min_code" json:"min_code"`                              // the smallest error code (an int)
	MaxCode              int                 `yaml:"max_code" json:"max_code"`                              // the biggest error code (an int)
	NextCode             int                 `yaml:"next_code" json:"next_code"`                            // the next error code to use, taken from ComponentInfo
	DuplicateCodes       map[string][]string `yaml:"duplicate_codes" json:"duplicate_codes"`                // duplicate error codes
	DuplicateNames       []string            `yaml:"duplicate_names" json:"duplicate_names"`                // duplicate error names
	CallExprCodes        []string            `yaml:"call_expr_codes" json:"call_expr_codes"`                // codes set by call expressions instead of literals
	IntCodes             []string            `yaml:"int_codes" json:"int_codes"`                            // all error codes as integers
	DeprecatedNewDefault []string            `yaml:"deprecated_new_default" json:"deprecated_new_default" ` // list of files with usage of deprecated NewDefault func
}
