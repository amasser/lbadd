package token

//go:generate stringer -type=Type
type Type uint16

const (
	Unknown Type = iota
	EOF
)
