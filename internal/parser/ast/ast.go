package ast

type Node interface {
	Positioner
	Lengther
	Typer
}

type Positioner interface {
	Line() int
	Col() int
	Offset() int
}

type Lengther interface {
	Length() int
}

type Typer interface {
	Type() NodeType
}
