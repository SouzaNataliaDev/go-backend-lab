package main

import (
	"context"
	"fmt"
	"time"
)

// context vem como primeiro parâmetro. Passe o contexto recebido adiante.
func lookup(ctx context.Context, duration time.Duration) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	timer := time.NewTimer(duration) // Simula I/O cancelável sem uma goroutine extra.
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-timer.C:
		return "cotação encontrada", nil
	}
}
func main() {
	// O contexto filho herda cancelamento e o prazo mais curto de seus ancestrais.
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel() // Libera recursos mesmo quando a operação termina antes do prazo.
	// Não há corrida com o timer de sucesso: duração zero usa prazo já vencido.
	expired, stop := context.WithTimeout(ctx, 0)
	defer stop()
	_, err := lookup(expired, time.Hour)
	fmt.Println("Prazo:", err)

	parent, cancelParent := context.WithCancel(context.Background())
	child, cancelChild := context.WithCancel(parent)
	defer cancelChild()
	cancelParent()
	_, err = lookup(child, time.Hour)
	fmt.Println("Cancelamento herdado:", err)

	value, err := lookup(context.Background(), 0)
	fmt.Println("Sucesso:", value, "Erro:", err)
	// USO REAL: http.NewRequestWithContext(ctx, ...) e db.QueryContext(ctx, ...).
	// Context sinaliza cancelamento; não mata goroutines nem força funções a parar.
	// Não substitua um contexto recebido por Background no meio da operação.
	// Não guarde ctx na struct, não passe nil e não use Value como opções gerais.
}
