package iterations

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func BenchmarkRepeatStd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatStd("a", 10)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("custom implementation", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})

	t.Run("standard implementation", func(t *testing.T) {
		repeated := RepeatStd("a", 5)
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})
}
