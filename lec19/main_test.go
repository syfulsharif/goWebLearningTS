package main

import "testing"

func TestDivide(t *testing.T) {
	_, err := divide(100.0, 10.0)

	if err != nil {
		t.Error("Got an error when we should not have!")
	}
}
func TestBadDivide(t *testing.T) {
	_, err := divide(100.0, 0)

	if err == nil {
		t.Error("Got no error when we should have!")
	}
}
