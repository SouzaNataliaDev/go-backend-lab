package workers

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
)

func TestPoolProcessesExactlyOnceWithinLimit(t *testing.T) {
	var active, maximum atomic.Int32
	seen := make([]atomic.Int32, 100)
	jobs := make([]int, 100)
	for i := range jobs {
		jobs[i] = i
	}
	err := Run(context.Background(), 3, jobs, func(ctx context.Context, id int) error {
		n := active.Add(1)
		defer active.Add(-1)
		for old := maximum.Load(); n > old; old = maximum.Load() {
			if maximum.CompareAndSwap(old, n) {
				break
			}
		}
		seen[id].Add(1)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if maximum.Load() > 3 {
		t.Fatal("concorrência excedida")
	}
	for i := range seen {
		if seen[i].Load() != 1 {
			t.Fatalf("job %d executado %d vezes", i, seen[i].Load())
		}
	}
}
func TestPoolStopsOnError(t *testing.T) {
	cause := errors.New("falha")
	if err := Run(context.Background(), 2, []int{1, 2, 3}, func(context.Context, int) error { return cause }); !errors.Is(err, cause) {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Run(ctx, 2, []int{1}, func(context.Context, int) error { t.Error("executou após cancelamento"); return nil }); !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
