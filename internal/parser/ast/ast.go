package ast

type Node interface {
	Positioner
	Typer
}

type Positioner interface {
	Line() int
	Col() int
	Offset() int
}

type Typer interface {
	Type() NodeType
}
