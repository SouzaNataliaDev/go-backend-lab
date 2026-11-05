package strategy

import (
	"context"
	"errors"
	"testing"
)

type failing struct{ err error }

func (f failing) Pay(context.Context, int64) (string, error) { return "", f.err }
func TestCheckout(t *testing.T) {
	for _, payment := range []Payment{Pix{}, Card{}} {
		c, err := NewCheckout(payment)
		if err != nil {
			t.Fatal(err)
		}
		if receipt, err := c.Finish(context.Background(), 100); err != nil || receipt == "" {
			t.Fatalf("receipt=%q err=%v", receipt, err)
		}
		if _, err := c.Finish(context.Background(), 0); err == nil {
			t.Fatal("aceitou zero")
		}
	}
	if _, err := NewCheckout(nil); err == nil {
		t.Fatal("aceitou nil")
	}
	cause := errors.New("falhou")
	c, _ := NewCheckout(failing{cause})
	if _, err := c.Finish(context.Background(), 100); !errors.Is(err, cause) {
		t.Fatalf("causa perdida: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := c.Finish(ctx, 100); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
