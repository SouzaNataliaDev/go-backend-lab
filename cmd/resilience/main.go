package main

import (
	"context"
	"errors"
	"github.com/SouzaNataliaDev/go-backend-lab/resilience"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	temporary := errors.New("provedor indisponível")
	attempts := 0
	err := resilience.Retry(ctx, 3, 10*time.Millisecond, 100*time.Millisecond, func(err error) bool { return errors.Is(err, temporary) }, func(ctx context.Context) error {
		attempts++
		if attempts < 3 {
			return temporary
		}
		return nil
	})
	log.Printf("retry: tentativas=%d erro=%v", attempts, err)
	breaker, err := resilience.NewBreaker(2, time.Second)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		err = breaker.Do(ctx, func(context.Context) error { return temporary })
		log.Printf("circuit breaker: chamada=%d erro=%v", i+1, err)
	}
}
