package main

import "testing"

func TestAvg(t *testing.T) {
	r, g, b := 10*256, 10*256, 10*256
	a := avg(uint32(r), uint32(g), uint32(b))

	if a != 10 {
		t.Errorf("wanted 10, got %d", a)
	}
}

func TestMapVal(t *testing.T) {
	v := 10
	g := mapVal(v, 0, 100, 0, 1000)

	if g != 100 {
		t.Errorf("wanted 100, got %d", g)
	}
}
