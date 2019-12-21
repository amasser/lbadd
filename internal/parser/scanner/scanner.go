package scanner

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

	line, lastCol, col int
}

func New(input []rune) *Scanner {
	return &Scanner{
		input: input,
		start: 0,
		pos:   0,

		current: initial,

		line:    1, // line starts at 1, because it should be human readable and editor line and column numbers usually start at 1
		lastCol: 1,
		col:     1, // col starts at 1, because it should be human readable and editor line and column numbers usually start at 1
	}
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

type matcher interface {
	Matches(rune) bool
}

func (s *Scanner) accept(m matcher) bool {
	if m.Matches(s.next()) {
		return true
	}
	s.goback()
	return false
}
