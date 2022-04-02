package arraysum

import (
	"reflect"
	"testing"
)

func assertSums(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v got %v", want, got)
	}
}

func assertEquals(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("expected '%d' got '%d'", want, got)
	}
}

func TestSum(t *testing.T) {
	t.Run("sum slice of 4 ints", func(t *testing.T) {
		result := SumArray([]int{2, 4, 6, 3})
		expected := 15
		assertEquals(t, result, expected)
	})
	t.Run("sum multiple arrays output single result", func(t *testing.T) {
		result := SumAllArrays([]int{1, 2}, []int{3, 4})
		expected := 10
		assertEquals(t, result, expected)
	})
	t.Run("sum multiple arrays output array", func(t *testing.T) {
		result := SumArrays([]int{4, 5}, []int{6, 7})
		expected := []int{9, 13}
		assertSums(t, result, expected)
	})
	t.Run("sum single array output array", func(t *testing.T) {
		result := SumArrays([]int{4, 5, 6})
		expected := []int{15}
		assertSums(t, result, expected)
	})
}

func TestSumTails(t *testing.T) {
	t.Run("sum tails valid", func(t *testing.T) {
		result := SumTails([]int{5, 6}, []int{1, 6, 8})
		expected := []int{6, 14}
		assertSums(t, result, expected)
	})
	t.Run("sum tails invalid", func(t *testing.T) {
		result := SumTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		assertSums(t, result, expected)
	})
}
