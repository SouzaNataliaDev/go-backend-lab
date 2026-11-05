package main

import "fmt"

func main() {
	// Um array tem tamanho fixo. Um slice descreve parte de um array subjacente.
	array := [3]int{10, 20, 30}
	values := array[:]
	window := values[:2] // len=2, cap=3: compartilha memória com values.
	window[0] = 99
	fmt.Println("Compartilhado:", values) // [99 20 30]

	// append pode reaproveitar o array quando existe capacidade disponível.
	window = append(window, 40)
	fmt.Println("Append com capacidade:", values) // [99 20 40]

	// Faça uma cópia quando precisar modificar sem afetar a origem.
	independent := make([]int, len(values))
	copy(independent, values)
	independent[0] = 1
	fmt.Println("Cópia:", independent, "Original:", values)

	// A expressão de três índices limita capacidade. O append aloca outro array.
	limited := values[:2:2]
	limited = append(limited, 50)
	fmt.Println("Capacidade limitada:", limited, "Original:", values)
	// Isso NÃO isola os dois elementos existentes: limited[0] ainda compartilhava
	// memória antes do append. Para isolamento completo, use copy.

	var items []string          // slice nil: len=0 e cap=0; append funciona.
	items = append(items, "Go") // Sempre guarde o slice retornado por append.
	fmt.Println("Lista:", items)
	// USO: resultados de consultas, lotes de mensagens e dados de uma API.
}
