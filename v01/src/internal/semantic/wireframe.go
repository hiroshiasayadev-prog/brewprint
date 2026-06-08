package semantic

type WireframeElement struct {
	Type        string
	ID          string
	Label       string
	Children    []WireframeElement
	Cols        int
	Fires       string
	Disabled    bool
	Placeholder string
	Span        int
	Layout      *WireframeLayout
}

type WireframeLayout struct {
	Width     any
	Height    any
	MinWidth  int
	MinHeight int
	Grow      bool
	Gap       int
	Padding   any
	Align     string
	Justify   string
	Scroll    string
}
