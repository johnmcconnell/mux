package intmux

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup(t *testing.T) (*Mux, *assert.Assertions) {
	a := assert.New(t)
	m := New(nil)

	return m, a
}

type testHandler1 struct {
}

func (t testHandler1) Handle(x int, y *int) {
	*y = x
}

type testHandler2 struct {
}

func (t testHandler2) Handle(x int, y *int) {
	*y = (2 * x)
}

func TestHasHandler(t *testing.T) {
	m, assert := setup(t)

	m.HandleWith(
		1,
		testHandler1{},
	)

	m.HandleWith(
		2,
		testHandler2{},
	)

	h := m.HasHandler(1)

	assert.True(
		h,
		"this mux has a handler at 1",
	)

	h = m.HasHandler(4)

	assert.False(
		h,
		"this mux does not have a handler at 10",
	)
}

func TestSelect(t *testing.T) {
	m, assert := setup(t)

	m.HandleWith(
		1,
		testHandler1{},
	)

	m.HandleWith(
		2,
		testHandler2{},
	)

	r := 0

	m.Select(1, 2, &r)

	assert.Equal(
		2,
		r,
		"result is 2",
	)

	m.Select(2, 5, &r)

	assert.Equal(
		10,
		r,
		"result is 2",
	)
}
