package main

import "testing"

func TestPalindrome(t *testing.T) {
	output := Palindrome(100, 150)
	expected := 5

	if output != expected {
		t.Errorf("expected %d but go %d", expected, output)
	}
}
