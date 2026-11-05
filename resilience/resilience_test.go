package resilience

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	temporary := errors.New("temporário")
	calls := 0
	err := Retry(context.Background(), 3, time.Nanosecond, time.Nanosecond, func(err error) bool { return errors.Is(err, temporary) }, func(context.Context) error {
		calls++
		if calls < 3 {
			return temporary
		}
		return nil
	})
	if err != nil || calls != 3 {
		t.Fatalf("calls=%d err=%v", calls, err)
	}
	calls = 0
	err = Retry(context.Background(), 3, time.Nanosecond, time.Nanosecond, func(error) bool { return false }, func(context.Context) error { calls++; return temporary })
	if !errors.Is(err, temporary) || calls != 1 {
		t.Fatalf("calls=%d err=%v", calls, err)
	}
}
func TestRetryCancellationDuringBackoff(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := Retry(ctx, 3, time.Hour, time.Hour, func(error) bool { return true }, func(context.Context) error { cancel(); return errors.New("falha") })
	if !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
func TestBreakerRecoveryAndSingleProbe(t *testing.T) {
	b, _ := NewBreaker(1, time.Second)
	now := time.Now()
	b.now = func() time.Time { return now }
	failure := errors.New("falha")
	if err := b.Do(context.Background(), func(context.Context) error { return failure }); !errors.Is(err, failure) {
		t.Fatal(err)
	}
	if err := b.Do(context.Background(), func(context.Context) error { t.Fatal("circuito deveria bloquear"); return nil }); !errors.Is(err, ErrCircuitOpen) {
		t.Fatal(err)
	}
	now = now.Add(time.Second)
	entered, release := make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = b.Do(context.Background(), func(context.Context) error { close(entered); <-release; return nil })
	}()
	<-entered
	if err := b.Do(context.Background(), func(context.Context) error { t.Error("segunda sonda executada"); return nil }); !errors.Is(err, ErrCircuitOpen) {
		t.Fatal(err)
	}
	close(release)
	wg.Wait()
	if err := b.Do(context.Background(), func(context.Context) error { return nil }); err != nil {
		t.Fatal(err)
	}
}
func TestStaleSuccessDoesNotCloseOpenCircuit(t *testing.T) {
	b, _ := NewBreaker(1, time.Second)
	entered, release := make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = b.Do(context.Background(), func(context.Context) error { close(entered); <-release; return nil })
	}()
	<-entered
	_ = b.Do(context.Background(), func(context.Context) error { return errors.New("falhou") })
	close(release)
	wg.Wait()
	if err := b.Do(context.Background(), func(context.Context) error { return nil }); !errors.Is(err, ErrCircuitOpen) {
		t.Fatalf("resposta antiga fechou circuito: %v", err)
	}
}
