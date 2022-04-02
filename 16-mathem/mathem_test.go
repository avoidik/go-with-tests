package mathem

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"math"
	"testing"
	"time"
)

//
// 360   x        360 * y
// --- = - => x = -------
//  60   y             60
//
//    360    180
// ------ = ----
// 2 * pi     pi
//
// (2 * pi) = 360 deg = 60 sec =  1 min
//                    3600 sec = 60 min = 1 h
//
// 2 * pi   y        pi * x
// ------ = - => y = ------ (sec)
//     60   x            30

func simpleTime(hh, mm, ss int) time.Time {
	return time.Date(2021, time.January, 1, hh, mm, ss, 0, time.UTC)
}

func TestHoursInRad(t *testing.T) {
	testCases := []struct {
		InputTime time.Time
		Angle     float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(21, 0, 0), math.Pi + math.Pi/2},
		{simpleTime(0, 1, 30), (90 * math.Pi) / (6 * 60 * 60)},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d hour %d min %d sec", testCase.InputTime.Hour(), testCase.InputTime.Minute(), testCase.InputTime.Second()), func(t *testing.T) {
			got := HoursInRad(testCase.InputTime)
			if !roughlyEqualFloat64(got, testCase.Angle) {
				t.Errorf("expected %f but got %f", testCase.Angle, got)
			}
		})
	}
}

func TestMinutesInRad(t *testing.T) {
	testCases := []struct {
		InputTime time.Time
		Angle     float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 7), (7.0 * math.Pi) / (60 * 30)},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 45, 0), math.Pi + math.Pi/2},
		{simpleTime(0, 45, 15), math.Pi + math.Pi/2 + (math.Pi*15)/(60*30)},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d min %d sec", testCase.InputTime.Minute(), testCase.InputTime.Second()), func(t *testing.T) {
			got := MinutesInRad(testCase.InputTime)
			if !roughlyEqualFloat64(got, testCase.Angle) {
				t.Errorf("expected %f but got %f", testCase.Angle, got)
			}
		})
	}
}

func TestSecondsInRad(t *testing.T) {
	testCases := []struct {
		InputTime time.Time
		Angle     float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 7), (7.0 * math.Pi) / 30},
		{simpleTime(0, 0, 15), math.Pi / 2},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), math.Pi + math.Pi/2},
		{simpleTime(0, 0, 60), 0},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d sec", testCase.InputTime.Second()), func(t *testing.T) {
			got := SecondsInRad(testCase.InputTime)
			if !roughlyEqualFloat64(got, testCase.Angle) {
				t.Errorf("expected %.2f but got %.2f", testCase.Angle, got)
			}
		})
	}

}

func TestSecondHandPoint(t *testing.T) {
	testCases := []struct {
		InputTime time.Time
		Point     Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d sec", testCase.InputTime.Second()), func(t *testing.T) {
			want := testCase.Point
			got := SecondHandPoint(testCase.InputTime)

			if !roughlyEqualPoint(t, got, want) {
				t.Errorf("expected %v but got %v", want, got)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	testCases := []struct {
		InputTime time.Time
		Point     Point
	}{
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d min %d sec", testCase.InputTime.Minute(), testCase.InputTime.Second()), func(t *testing.T) {
			want := testCase.Point
			got := MinuteHandPoint(testCase.InputTime)

			if !roughlyEqualPoint(t, got, want) {
				t.Errorf("expected %v but got %v", want, got)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	testCases := []struct {
		InputTime time.Time
		Point     Point
	}{
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(21, 0, 0), Point{-1, 0}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d hour %d min %d sec", testCase.InputTime.Hour(), testCase.InputTime.Minute(), testCase.InputTime.Second()), func(t *testing.T) {
			want := testCase.Point
			got := HourHandPoint(testCase.InputTime)

			if !roughlyEqualPoint(t, got, want) {
				t.Errorf("expected %v but got %v", want, got)
			}
		})
	}
}

func TestSVGWriterSecondHand(t *testing.T) {
	testCases := []struct {
		tm   time.Time
		line Line
	}{
		{simpleTime(00, 00, 00), Line{150, 150, 150, 60}},
		{simpleTime(00, 00, 30), Line{150, 150, 150, 240}},
		{simpleTime(00, 00, 45), Line{150, 150, 60, 150}},
	}

	for _, testCase := range testCases {
		b := bytes.Buffer{}
		SVGWriter(&b, testCase.tm)

		svg := SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		want := testCase.line

		if !containsLine(want, svg.Line) {
			t.Errorf("Expected to find the second hand %+v, in the SVG output %+v", want, svg.Line)
		}
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	testCases := []struct {
		tm   time.Time
		line Line
	}{
		{simpleTime(00, 00, 00), Line{150, 150, 150, 60}},
		{simpleTime(00, 30, 00), Line{150, 150, 150, 240}},
		{simpleTime(00, 45, 00), Line{150, 150, 60, 150}},
	}

	for _, testCase := range testCases {
		b := bytes.Buffer{}
		SVGWriter(&b, testCase.tm)

		svg := SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		want := testCase.line

		if !containsLine(want, svg.Line) {
			t.Errorf("Expected to find the minute hand %+v, in the SVG output %+v", want, svg.Line)
		}
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	testCases := []struct {
		tm   time.Time
		line Line
	}{
		{simpleTime(00, 00, 00), Line{150, 150, 150, 100}},
		{simpleTime(6, 00, 00), Line{150, 150, 150, 200}},
	}

	for _, testCase := range testCases {
		b := bytes.Buffer{}
		SVGWriter(&b, testCase.tm)

		svg := SVG{}
		xml.Unmarshal(b.Bytes(), &svg)

		want := testCase.line

		if !containsLine(want, svg.Line) {
			t.Errorf("Expected to find the minute hand %+v, in the SVG output %+v", want, svg.Line)
		}
	}
}

func containsLine(find Line, lines []Line) bool {
	for _, line := range lines {
		if line == find {
			return true
		}
	}
	return false
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(t *testing.T, a, b Point) bool {
	t.Helper()

	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
