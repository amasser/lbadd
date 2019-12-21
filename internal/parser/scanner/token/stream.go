package token

type Stream struct {
	ch chan Token

	peeked *Token

	closed bool
}

func NewStream() *Stream {
	return &Stream{
		ch:     make(chan Token, 5),
		peeked: nil,
		closed: false,
	}
}

// Peek returns the first element of the stream, waiting for it to become
// available if necessary. However, it does NOT remove the element from the
// stream.
func (s *Stream) Peek() Token {
	s.ensureNotClosed()
	if s.peeked != nil {
		return *s.peeked
	}
	t := s.Take()
	s.peeked = &t
	return t
}

// Take returns the first element of the stream, waiting for it to become
// available if necessary. Take also removes the element from the stream.
func (s *Stream) Take() Token {
	s.ensureNotClosed()
	if s.peeked != nil {
		t := *s.peeked
		s.peeked = nil
		return t
	}
	return <-s.ch
}

// Push pushes a token onto the stream, waiting if the stream is full.
func (s *Stream) Push(t Token) {
	s.ensureNotClosed()
	s.ch <- t
}

func (s Stream) ensureNotClosed() {
	if s.closed {
		panic("token stream is closed")
	}
}

func (s *Stream) Close() error {
	s.closed = true
	close(s.ch)
	return nil
}

func (s Stream) Closed() bool {
	return s.closed
}
