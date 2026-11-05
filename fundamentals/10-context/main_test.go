package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestLookupCancellationPropagates(t *testing.T) {
	parent, cancel := context.WithCancel(context.Background())
	child, stop := context.WithCancel(parent)
	defer stop()
	cancel()
	if _, err := lookup(child, time.Hour); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
func TestLookupDeadlineWhileWaiting(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	if _, err := lookup(ctx, time.Hour); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal(err)
	}
}
