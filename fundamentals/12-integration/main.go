package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Quote struct {
	Provider string
	Cents    int64
}
type Result struct {
	Quote Quote
	Err   error
}
type Provider interface {
	Quote(context.Context) (Quote, error)
}
type Simulated struct {
	Name  string
	Cents int64
	Delay time.Duration
}

func (p Simulated) Quote(ctx context.Context) (Quote, error) {
	timer := time.NewTimer(p.Delay)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return Quote{}, ctx.Err()
	case <-timer.C:
		return Quote{Provider: p.Name, Cents: p.Cents}, nil
	}
}

// collect dispara um número limitado de consultas e reúne resultados (fan-out/fan-in).
// Para um volume arbitrário, use worker pool. Os provedores devem respeitar ctx.
func collect(ctx context.Context, providers []Provider) <-chan Result {
	// Buffer comporta uma resposta por provedor, mesmo se o consumidor sair cedo.
	results := make(chan Result, len(providers))
	var wg sync.WaitGroup
	for _, provider := range providers {
		wg.Add(1)
		go func(p Provider) {
			defer wg.Done()
			quote, err := p.Quote(ctx)
			if err != nil {
				err = fmt.Errorf("consultar provedor: %w", err)
			}
			results <- Result{Quote: quote, Err: err}
		}(provider)
	}
	// Só o coordenador fecha, após todos os envios terminarem.
	go func() { wg.Wait(); close(results) }()
	return results
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	providers := []Provider{
		Simulated{Name: "A", Cents: 10000, Delay: time.Millisecond},
		Simulated{Name: "B", Cents: 9900, Delay: 2 * time.Millisecond},
	}
	var best *Quote
	for result := range collect(ctx, providers) {
		if result.Err != nil {
			fmt.Println("Falha:", result.Err)
			continue
		}
		if best == nil || result.Quote.Cents < best.Cents {
			quote := result.Quote // cópia própria do resultado escolhido.
			best = &quote
		}
	}
	if best == nil {
		fmt.Println("Nenhuma cotação disponível")
		return
	}
	fmt.Printf("Melhor cotação: %s, %d centavos\n", best.Provider, best.Cents)
	// É uma simulação local: não faz requisições externas nem pagamentos.
	// A quantidade de goroutines e o buffer crescem com a lista de provedores.
	// Cancelar sinaliza parada, mas coletar ainda depende de provedores cooperativos.
}
