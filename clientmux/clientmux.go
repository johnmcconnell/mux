package clientmux

import c "io"

// Handler handles commands
type Handler interface {
	Handle(in []string, out c.Writer)
}

// New returns a new handler
func New(h map[string]Handler) *Mux {
	if h == nil {
		h = map[string]Handler{}
	}

	m := Mux{
		Handlers: h,
	}

	return &m
}

// Mux handles choosing handlers for the given command
// in the request
type Mux struct {
	Handlers map[string]Handler
}

// HandleWith sets the handler for a given command
func (m *Mux) HandleWith(s string, h Handler) {
	m.Handlers[s] = h
}

// HasHandler checks that the mux has a handler for
// that string
func (m *Mux) HasHandler(s string) bool {
	_, ok := m.Handlers[s]

	return ok
}

// Select tell the mux to select and run the command
func (m *Mux) Select(s string, in []string, out c.Writer) bool {
	h, ok := m.Handlers[s]

	if ok {
		h.Handle(in, out)
		return true
	}

	return false
}

// Here we will test that the types parameters are ok...
func testTypes(arg0 string, arg1 []string, arg2 c.Writer) {
	f := func(interface{}, interface{}, interface{}) {} // this func does nothing...
	f(arg0, arg1, arg2)
}
