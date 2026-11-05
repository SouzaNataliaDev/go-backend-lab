package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

type broken struct{ cause error }

func (p broken) Quote(context.Context) (Quote, error) { return Quote{}, p.cause }
func TestCollectPartialFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cause := errors.New("provedor falhou")
	success, failures := 0, 0
	for r := range collect(ctx, []Provider{Simulated{Name: "A", Cents: 100}, broken{cause}}) {
		if errors.Is(r.Err, cause) {
			failures++
		} else if r.Err == nil && r.Quote.Cents == 100 {
			success++
		} else {
			t.Fatalf("resultado inesperado: %+v", r)
		}
	}
	if success != 1 || failures != 1 {
		t.Fatalf("sucessos=%d falhas=%d", success, failures)
	}
}
func TestCollectCancelledAndCloses(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	// Timeout é só proteção do teste contra uma regressão que vaze goroutines.
	watchdog := time.NewTimer(time.Second)
	defer watchdog.Stop()
	results := collect(ctx, []Provider{Simulated{Name: "lento", Delay: time.Hour}})
	select {
	case r := <-results:
		if !errors.Is(r.Err, context.Canceled) {
			t.Fatal(r.Err)
		}
	case <-watchdog.C:
		t.Fatal("cancelamento não encerrou a consulta")
	}
	select {
	case _, ok := <-results:
		if ok {
			t.Fatal("resultado extra")
		}
	case <-watchdog.C:
		t.Fatal("canal não foi fechado")
	}
}
func TestCollectEmpty(t *testing.T) {
	for range collect(context.Background(), nil) {
		t.Fatal("resultado inesperado")
	}
}
