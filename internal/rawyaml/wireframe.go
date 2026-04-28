package rawyaml

type WireframeElement struct {
	Type        string             `yaml:"type"`
	ID          string             `yaml:"id"`
	Label       string             `yaml:"label"`
	Children    []WireframeElement `yaml:"children"`
	Cols        int                `yaml:"cols"`
	Fires       string             `yaml:"fires"`
	Disabled    bool               `yaml:"disabled"`
	Placeholder string             `yaml:"placeholder"`
	Span        int                `yaml:"span"`
	Layout      *WireframeLayout   `yaml:"layout"`
}

type WireframeLayout struct {
	Width     any    `yaml:"width"`
	Height    any    `yaml:"height"`
	MinWidth  int    `yaml:"min_width"`
	MinHeight int    `yaml:"min_height"`
	Grow      bool   `yaml:"grow"`
	Gap       int    `yaml:"gap"`
	Padding   any    `yaml:"padding"`
	Align     string `yaml:"align"`
	Justify   string `yaml:"justify"`
	Scroll    string `yaml:"scroll"`
}
