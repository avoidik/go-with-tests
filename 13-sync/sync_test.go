package synccount

import (
	"sync"
	"testing"
)

func assertCount(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("expected %d got %d", want, got.Value())
	}
}

func TestCounter(t *testing.T) {
	t.Run("increment by 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCount(t, counter, 3)
	})
	t.Run("concurrent", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(t, counter, wantedCount)
	})
}
