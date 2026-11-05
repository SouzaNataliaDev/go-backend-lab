// Package resilience reúne mecanismos didáticos para controlar falhas.
package resilience

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

// Retry repete somente erros classificados como transitórios, com backoff e jitter.
// Use apenas com operações idempotentes. O contexto deve ter prazo total.
func Retry(ctx context.Context, attempts int, base, max time.Duration, transient func(error) bool, operation func(context.Context) error) error {
	if attempts < 1 || base <= 0 || max < base || transient == nil || operation == nil {
		return errors.New("configuração de retry inválida")
	}
	delay := base
	for attempt := 0; attempt < attempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		err := operation(ctx)
		if err == nil {
			return nil
		}
		if attempt == attempts-1 || !transient(err) {
			return err
		}
		// Equal jitter: espera entre metade do backoff e o backoff completo.
		wait := delay/2 + time.Duration(rand.Int63n(int64(delay-delay/2)))
		timer := time.NewTimer(wait)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
		if delay > max/2 {
			delay = max
		} else {
			delay *= 2
		}
	}
	return nil
}
