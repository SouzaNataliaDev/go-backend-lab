package resilience

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrCircuitOpen = errors.New("circuito aberto")

// Breaker conta falhas consecutivas. Após cooldown, permite uma única sonda.
// generation impede respostas antigas de fecharem um circuito recém-aberto.
type Breaker struct {
	mu                  sync.Mutex
	threshold, failures int
	cooldown            time.Duration
	openedAt            time.Time
	open, probing       bool
	generation          uint64
	now                 func() time.Time
}

func NewBreaker(threshold int, cooldown time.Duration) (*Breaker, error) {
	if threshold < 1 || cooldown <= 0 {
		return nil, errors.New("configuração de circuito inválida")
	}
	return &Breaker{threshold: threshold, cooldown: cooldown, now: time.Now}, nil
}
func (b *Breaker) Do(ctx context.Context, operation func(context.Context) error) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	b.mu.Lock()
	probe := false
	if b.open {
		if b.probing || b.now().Sub(b.openedAt) < b.cooldown {
			b.mu.Unlock()
			return ErrCircuitOpen
		}
		b.probing = true
		probe = true
	}
	generation := b.generation
	b.mu.Unlock()
	err := operation(ctx)
	b.mu.Lock()
	defer b.mu.Unlock()
	if generation != b.generation {
		return err
	}
	if probe {
		b.probing = false
		b.generation++
		if err == nil {
			b.open = false
			b.failures = 0
		} else {
			b.openedAt = b.now()
		}
		return err
	}
	if err == nil {
		b.failures = 0
		return nil
	}
	b.failures++
	if b.failures >= b.threshold {
		b.open = true
		b.openedAt = b.now()
		b.generation++
	}
	return err
}
