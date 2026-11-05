package main

import "testing"

func TestConcurrentCounter(t *testing.T) {
	if got := count(20, 1000); got != 20000 {
		t.Fatalf("perdeu incrementos: %d", got)
	}
}
