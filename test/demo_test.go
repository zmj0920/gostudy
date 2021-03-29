package demo1

import "testing"

const (
	Monday = 1 + iota
	Tuesday
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}
