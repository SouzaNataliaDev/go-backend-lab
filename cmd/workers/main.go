package main

import (
	"context"
	"github.com/SouzaNataliaDev/go-backend-lab/scalability/workers"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := workers.Run(ctx, 3, []int{1, 2, 3, 4, 5, 6}, func(ctx context.Context, id int) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(10 * time.Millisecond):
		}
		log.Printf("job %d concluído", id)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
