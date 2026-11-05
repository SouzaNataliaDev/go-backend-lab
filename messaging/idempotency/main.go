// Exemplo local de consumidor idempotente. Não conecta a um broker.
package main

import (
	"errors"
	"fmt"
	"sync"
)

type Event struct {
	ID    string
	Cents int64
}

// Consumer não deve ser copiado após uso: contém um mutex.
// Estado e deduplicação ficam juntos para atualizar de forma atômica neste processo.
type Consumer struct {
	mu        sync.Mutex
	processed map[string]int64
	balance   int64
}

func NewConsumer() *Consumer { return &Consumer{processed: make(map[string]int64)} }

var ErrKeyConflict = errors.New("mesmo ID recebido com conteúdo diferente")

// Apply devolve applied=false para uma repetição que já foi processada.
func (c *Consumer) Apply(event Event) (applied bool, err error) {
	if event.ID == "" || event.Cents <= 0 {
		return false, errors.New("evento inválido")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if previous, ok := c.processed[event.ID]; ok {
		if previous != event.Cents {
			return false, ErrKeyConflict
		}
		return false, nil
	}
	// O efeito e o registro ficam na mesma seção crítica.
	// Não coloque uma chamada externa aqui: mutex não cria transação distribuída.
	c.balance += event.Cents
	c.processed[event.ID] = event.Cents
	return true, nil
}
func (c *Consumer) Balance() int64 { c.mu.Lock(); defer c.mu.Unlock(); return c.balance }
func main() {
	consumer := NewConsumer()
	event := Event{ID: "evt-42", Cents: 5000}
	for i := 1; i <= 2; i++ {
		applied, err := consumer.Apply(event)
		fmt.Printf("Entrega %d: aplicado=%v erro=%v\n", i, applied, err)
	}
	fmt.Println("Saldo em centavos:", consumer.Balance())
	// USO: evitar repetir o efeito de uma entrega duplicada.
	// LIMITES: memória é perdida no restart, não compartilhada entre réplicas,
	// cresce sem limite e não protege efeitos externos. Veja README e inbox.sql.
}
