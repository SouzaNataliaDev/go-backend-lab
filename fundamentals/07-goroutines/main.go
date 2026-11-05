package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []string{"estoque", "preço", "frete"}
	var wg sync.WaitGroup
	for _, task := range tasks {
		// Add antes de iniciar a goroutine, para Wait não retornar cedo demais.
		wg.Add(1)
		go func(name string) {
			defer wg.Done() // Sinaliza a conclusão mesmo em retornos antecipados.
			fmt.Println("Consultando:", name)
			// Simulação de uma consulta independente.
		}(task) // Parâmetro explícito torna a captura do valor fácil de entender.
	}
	wg.Wait() // Espera todas as tarefas. Não usamos Sleep como sincronização.
	fmt.Println("Todas as consultas concluídas")
	// USO: chamadas de I/O independentes. Uma goroutine é gerenciada pelo runtime;
	// ela não corresponde necessariamente a uma thread exclusiva do sistema.
	// Concorrência é sobre tarefas em progresso; paralelismo é execução simultânea.
	// A ordem das três primeiras linhas pode mudar a cada execução.
	// WaitGroup espera tarefas; não protege variáveis compartilhadas e não retorna erros.
	// Limite a concorrência para entradas grandes: veja scalability/workers.
}
