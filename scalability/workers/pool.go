// Package workers demonstra concorrência limitada e backpressure.
package workers

import (
	"context"
	"errors"
	"sync"
)

// Run usa uma fila sem buffer: produtor espera um worker ficar disponível.
// fn deve respeitar o contexto e ser segura para execução concorrente.
// Na primeira falha, cancela as tarefas restantes e aguarda todos os workers.
func Run(ctx context.Context, count int, jobs []int, fn func(context.Context, int) error) error {
	if count < 1 || fn == nil {
		return errors.New("configuração de workers inválida")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	queue := make(chan int)
	var wg sync.WaitGroup
	var once sync.Once
	var first error
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-queue:
					if !ok {
						return
					}
					if ctx.Err() != nil {
						return
					}
					if err := fn(ctx, job); err != nil {
						once.Do(func() { first = err; cancel() })
						return
					}
				}
			}
		}()
	}
feed:
	for _, job := range jobs {
		select {
		case <-ctx.Done():
			break feed
		case queue <- job:
		}
	}
	close(queue)
	wg.Wait()
	if first != nil {
		return first
	}
	return ctx.Err()
}
