// Package decorator adiciona observabilidade sem alterar a estratégia original.
package decorator

import (
	"context"
	"github.com/SouzaNataliaDev/go-backend-lab/patterns/strategy"
	"log/slog"
	"time"
)

// Observed exige dependências não nulas. Não registre dados de cartão ou recibos.
type Observed struct {
	Next   strategy.Payment
	Logger *slog.Logger
}

func (p Observed) Pay(ctx context.Context, cents int64) (string, error) {
	start := time.Now()
	receipt, err := p.Next.Pay(ctx, cents)
	p.Logger.InfoContext(ctx, "pagamento concluído", "duration", time.Since(start), "success", err == nil)
	return receipt, err
}
