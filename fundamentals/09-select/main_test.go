package main

import (
	"context"
	"errors"
	"testing"
)

func TestReceiveCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := receive(ctx, make(chan string)); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
func TestReceiveClosed(t *testing.T) {
	ch := make(chan string)
	close(ch)
	if value, err := receive(context.Background(), ch); err != nil || value != "sem mais resultados" {
		t.Fatalf("%q %v", value, err)
	}
}
