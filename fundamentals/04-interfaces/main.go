package main

import (
	"context"
	"fmt"
)

// O consumidor declara só o comportamento de que precisa.
// Não existe implements: um tipo satisfaz a interface pelos seus métodos.
type Notifier interface {
	Send(context.Context, string) error
}

type Console struct{}

func (Console) Send(ctx context.Context, message string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	fmt.Println("Notificação:", message)
	return nil
}

type Recorder struct{ Messages []string }

// Receiver ponteiro permite alterar a struct original.
func (r *Recorder) Send(ctx context.Context, message string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	r.Messages = append(r.Messages, message)
	return nil
}

// Verificação em compilação: só *Recorder possui o método com receiver ponteiro.
var _ Notifier = (*Recorder)(nil)

func notifyOrder(ctx context.Context, notifier Notifier) error {
	return notifier.Send(ctx, "Pedido aprovado")
}

func main() {
	if err := notifyOrder(context.Background(), Console{}); err != nil {
		fmt.Println(err)
	}
	recorder := &Recorder{}
	if err := notifyOrder(context.Background(), recorder); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Fake registrou:", recorder.Messages)

	// Uma interface contém tipo dinâmico + valor dinâmico.
	var pointer *Recorder
	var notifier Notifier = pointer
	fmt.Println("Ponteiro nil:", pointer == nil)   // true
	fmt.Println("Interface nil:", notifier == nil) // false: tem tipo *Recorder.
	// Não chame Send neste notifier: o método acima não aceita receiver nil.
	// USO: provedores de pagamento, repositórios e clientes externos substituíveis.
}
