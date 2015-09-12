package intmux

// Handler handles commands
type Handler interface {
	Handle(in int, out *int)
}

// New returns a new handler
func New(h map[int]Handler) *Mux {
	if h == nil {
		h = map[int]Handler{}
	}

	m := Mux{
		Handlers: h,
	}

	return &m
}

// Mux handles choosing handlers for the given command
// in the request
type Mux struct {
	Handlers map[int]Handler
}

// HandleWith sets the handler for a given command
func (m *Mux) HandleWith(s int, h Handler) {
	m.Handlers[s] = h
}

// HasHandler checks that the mux has a handler for
// that string
func (m *Mux) HasHandler(s int) bool {
	_, ok := m.Handlers[s]

	return ok
}

// Select tell the mux to select and run the command
func (m *Mux) Select(s int, in int, out *int) bool {
	h, ok := m.Handlers[s]

	if ok {
		h.Handle(in, out)
		return true
	}

	return false
}

// Here we will test that the types parameters are ok...
func testTypes(arg0 int, arg1 int, arg2 *int) {
	f := func(byte, interface{}, interface{}) {} // this func does nothing...
	f(byte(arg0), arg1, arg2)
}
