package qikmux

import b "io"
import c "io"

// Handler handles commands
type Handler interface {
	Handle(in b.Reader, out c.Writer)
}

// New returns a new handler
func New(h map[byte]Handler) *Mux {
	if h == nil {
		h = map[byte]Handler{}
	}

	m := Mux{
		Handlers: h,
	}

	return &m
}

// Mux handles choosing handlers for the given command
// in the request
type Mux struct {
	Handlers map[byte]Handler
}

// HandleWith sets the handler for a given command
func (m *Mux) HandleWith(s byte, h Handler) {
	m.Handlers[s] = h
}

// HasHandler checks that the mux has a handler for
// that string
func (m *Mux) HasHandler(s byte) bool {
	_, ok := m.Handlers[s]

	return ok
}

// Select tell the mux to select and run the command
func (m *Mux) Select(s byte, in b.Reader, out c.Writer) bool {
	h, ok := m.Handlers[s]

	if ok {
		h.Handle(in, out)
		return true
	}

	return false
}

// Here we will test that the types parameters are ok...
func testTypes(arg0 byte, arg1 b.Reader, arg2 c.Writer) {
	f := func(byte, interface{}, interface{}) {} // this func does nothing...
	f(byte(arg0), arg1, arg2)
}
