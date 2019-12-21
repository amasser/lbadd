package parser

import "github.com/tomarrell/lbadd/internal/parser/ast"

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(sql string) ast.Query {
	panic("TODO")
}
