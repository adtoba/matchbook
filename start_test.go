package main

import "testing"

func TestStart(t *testing.T) {
	start := "letsgo"

	if start != "letsgo" {
		t.Fatalf("test failed")
	}
}
