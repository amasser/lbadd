package scanner

import (
	"reflect"
	"runtime"

	"github.com/tomarrell/lbadd/internal/parser/scanner/matcher"
	"github.com/tomarrell/lbadd/internal/parser/scanner/token"
)

// state is a type alias for a function that takes a scanner and returns another
// state. Such functions (or states) will be invoked by the scanner. It will
// pass itself as the argument, and the returned state will be chained and
// executed next.
type state func(*Scanner) state

type Scanner struct {
	input []rune
	start int
	pos   int

	current state
	stream  *token.Stream

	line, lastCol, col int
}

func New(input []rune, stream *token.Stream) *Scanner {
	return &Scanner{
		input: input,
		start: 0,
		pos:   0,

		current: initial,
		stream:  stream,

		line:    1, // line starts at 1, because it should be human readable and editor line and column numbers usually start at 1
		lastCol: 1,
		col:     1, // col starts at 1, because it should be human readable and editor line and column numbers usually start at 1
	}
}

func (s *Scanner) Scan() {
	for !s.done() {
		nextState := s.current(s)
		if nextState == nil {
			panic("state '" + runtime.FuncForPC(reflect.ValueOf(s.current).Pointer()).Name() + "' evaluated to nil")
		}
		s.current = nextState
	}
	s.emit(token.EOF)
}

func (s *Scanner) done() bool {
	return s.pos >= len(s.input)
}

func (s *Scanner) next() rune {
	next := s.input[s.pos]

	// update line and column information
	if next == '\n' { // TODO if next is line delimiter
		s.line++
		s.lastCol = s.col
		s.col = 1
	} else {
		s.col++
	}

	return next
}

func (s *Scanner) goback() {
	s.pos--

	// update line and column information
	if s.col == 1 {
		s.line--
		s.col = s.lastCol
	} else {
		s.col--
	}
}

func (s *Scanner) emit(t token.Type) {
	tok := token.New(s.line, s.col, s.start, s.pos-s.start, t, string(s.input[s.start:s.pos]))
	s.stream.Push(tok)
}

func (s *Scanner) accept(m matcher.M) bool {
	if m.Matches(s.next()) {
		return true
	}
	s.goback()
	return false
}
func (s *Scanner) acceptSequence(seq string) bool {
	checkpoint := s.pos
	for _, r := range seq {
		if r != s.next() {
			s.pos = checkpoint
			return false
		}
	}
	return true
}
