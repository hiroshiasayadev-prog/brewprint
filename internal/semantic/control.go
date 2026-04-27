package semantic

type Branch struct {
	BaseNode
	Params []Param
}

type Fork struct {
	BaseNode
	Params []Param
}

type Join struct {
	BaseNode
	Params  []Param
	Returns *Return
}
