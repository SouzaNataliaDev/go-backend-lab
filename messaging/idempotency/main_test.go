package main

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
)

func TestConcurrentDuplicatesApplyOnce(t *testing.T) {
	consumer := NewConsumer()
	var applied atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ok, err := consumer.Apply(Event{ID: "same", Cents: 5000})
			if err != nil {
				t.Error(err)
			}
			if ok {
				applied.Add(1)
			}
		}()
	}
	wg.Wait()
	if applied.Load() != 1 || consumer.Balance() != 5000 {
		t.Fatalf("applied=%d balance=%d", applied.Load(), consumer.Balance())
	}
}
func TestKeyConflictAndValidation(t *testing.T) {
	c := NewConsumer()
	_, _ = c.Apply(Event{ID: "one", Cents: 100})
	if _, err := c.Apply(Event{ID: "one", Cents: 200}); !errors.Is(err, ErrKeyConflict) {
		t.Fatal(err)
	}
	if _, err := c.Apply(Event{Cents: 100}); err == nil {
		t.Fatal("aceitou ID vazio")
	}
	if c.Balance() != 100 {
		t.Fatal("evento inválido alterou saldo")
	}
}
