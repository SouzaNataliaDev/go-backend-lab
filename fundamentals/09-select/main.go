package main

import (
	"context"
	"fmt"
)

func receive(ctx context.Context, results <-chan string) (string, error) {
	select {
	case value, ok := <-results:
		if !ok {
			return "sem mais resultados", nil
		}
		return value, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
	// Se vários cases puderem avançar, select escolhe um pseudoaleatoriamente.
	// A posição dos cases não dá prioridade ao cancelamento.
}
func main() {
	ready := make(chan string, 1)
	ready <- "cotação recebida"
	value, err := receive(context.Background(), ready)
	fmt.Println("Resultado:", value, "Erro:", err)

	// Canal sem produtor: só o cancelamento poderá liberar este select.
	empty := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = receive(ctx, empty)
	fmt.Println("Cancelado:", err)

	select {
	case value := <-empty:
		fmt.Println(value)
	default:
		fmt.Println("Nenhum resultado disponível agora")
	}
	// USO: aguardar resultado ou prazo de uma requisição.
	// default torna o select não bloqueante. Dentro de loop sem espera,
	// ele pode criar busy loop e consumir CPU. Não é um timeout.
}
