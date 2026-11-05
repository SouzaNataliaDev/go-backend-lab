package main

import (
	"fmt"
	"sync"
)

// Counter deve ser usado por ponteiro: não copie structs com mutex após o uso.
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	// value++ é leitura + cálculo + escrita. Sem mutex haveria data race.
	c.value++
}
func (c *Counter) Value() int {
	// Leituras também precisam participar da sincronização se houver escritores.
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
func count(workers, each int) int {
	var counter Counter
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < each; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	return counter.Value()
}
func main() {
	fmt.Println("Total:", count(10, 1000))
	// USO: estado compartilhado local. Proteja maps da mesma forma quando necessário.
	// WaitGroup controla conclusão; Mutex controla acesso à seção crítica.
	// Evite segurar o mutex durante chamada HTTP ou outra operação lenta.
	// Rode: go run -race ./fundamentals/11-concurrency
	// Mutex não torna estado local compartilhado entre réplicas diferentes.
}
