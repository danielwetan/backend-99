package main

import "testing"

func TestSoal2(t *testing.T) {
	output := Solution("101102103104105106107108109111112113")
	expected := 110

	if output != expected {
		t.Errorf("expected %d but go %d", expected, output)
	}
}
