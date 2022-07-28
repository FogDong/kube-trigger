package api

type Source struct {
	Kind       string `json:"kind" yaml:"kind"`
	ApiVersion string `json:"apiVersion" yaml:"apiVersion"`
	Namespace  string `json:"namespace" yaml:"namespace"`
}

type Label struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}
type Filter struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
}
type App struct {
	Namespace string  `json:"namespace" yaml:"namespace"`
	Name      string  `json:"name" yaml:"name"`
	Labels    []Label `json:"labels" yaml:"labels"`
}
type Trigger struct {
	WatchKind Source   `json:"watch" yaml:"watch_kind"`
	Filters   []Filter `json:"filters" yaml:"filters"`
	To        App      `json:"to" yaml:"to"`
}
