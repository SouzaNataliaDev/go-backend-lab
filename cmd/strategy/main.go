package main

import (
	"context"
	"github.com/SouzaNataliaDev/go-backend-lab/patterns/decorator"
	"github.com/SouzaNataliaDev/go-backend-lab/patterns/strategy"
	"log"
	"log/slog"
)

func main() {
	for _, payment := range []strategy.Payment{strategy.Pix{}, strategy.Card{}} {
		checkout, err := strategy.NewCheckout(decorator.Observed{Next: payment, Logger: slog.Default()})
		if err != nil {
			log.Fatal(err)
		}
		receipt, err := checkout.Finish(context.Background(), 10000)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(receipt)
	}
}
