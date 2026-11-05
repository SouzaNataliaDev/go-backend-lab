// Package strategy demonstra troca de comportamento por uma interface pequena.
package strategy

import (
	"context"
	"errors"
	"fmt"
)

// Payment é definido pelo consumidor. Valores monetários são centavos inteiros.
type Payment interface {
	Pay(context.Context, int64) (string, error)
}

type Checkout struct{ payment Payment }

func NewCheckout(payment Payment) (*Checkout, error) {
	if payment == nil {
		return nil, errors.New("pagamento obrigatório")
	}
	return &Checkout{payment: payment}, nil
}
func (c *Checkout) Finish(ctx context.Context, cents int64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	if cents <= 0 {
		return "", errors.New("valor deve ser positivo")
	}
	receipt, err := c.payment.Pay(ctx, cents)
	if err != nil {
		return "", fmt.Errorf("finalizar checkout: %w", err)
	}
	return receipt, nil
}

// Pix e Card simulam provedores: nenhum pagamento real é feito.
type Pix struct{}

func (Pix) Pay(ctx context.Context, cents int64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	return fmt.Sprintf("Pix aprovado: %d centavos", cents), nil
}

type Card struct{}

func (Card) Pay(ctx context.Context, cents int64) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	return fmt.Sprintf("Cartão aprovado: %d centavos", cents), nil
}
