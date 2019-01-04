package types

// Knative represents simple Knative object
type Knative struct {
	Meta    *RequestMetadata       `json:"meta"`
	InfoURI string                 `json:"infoURI"`
	EnvVars map[string]interface{} `json:"envs"`
	Access  []*FsAccessGroup       `json:"access"`
}

// FsAccessGroup represents a collection of Fs access
type FsAccessGroup struct {
	Group string          `json:"group"`
	List  []*FsAccessInfo `json:"list"`
}

// FsAccessInfo represents the access assertion info
type FsAccessInfo struct {
	Path     string `json:"path"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
	Comment  string `json:"comment"`
}
