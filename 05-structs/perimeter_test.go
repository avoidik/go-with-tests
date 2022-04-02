package perimeter

import (
	"testing"
)

func assertArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("%T: expected %g got %g", shape, want, got)
	}
}

func assertPerimeter(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()
	if got != want {
		t.Errorf("%T: expected %g got %g", shape, want, got)
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		desc  string
		shape Shape
		want  float64
	}{
		{"calc rectangle", Rectangle{10.0, 10.0}, 40.0},
		{"calc circle", Circle{10.0}, 31.41592653589793},
		{"calc triangle", Triangle{5.0, 4.0, 3.0}, 12.0},
	}

	for _, testCase := range perimeterTests {
		t.Run(testCase.desc, func(t *testing.T) {
			assertPerimeter(t, testCase.shape, testCase.want)
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		desc  string
		shape Shape
		want  float64
	}{
		{"calc rectangle", Rectangle{7.0, 6.0}, 42.0},
		{"calc circle", Circle{10}, 314.1592653589793},
		{"calc triangle", Triangle{5.0, 4.0, 3.0}, 6.0},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.desc, func(t *testing.T) {
			assertArea(t, testCase.shape, testCase.want)
		})
	}
}
