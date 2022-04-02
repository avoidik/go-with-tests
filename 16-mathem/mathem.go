package mathem

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	hourHandLength   = 50.0
	minuteHandLength = 90.0
	secondHandLength = 90.0
	clockCenterX     = 150.0
	clockCenterY     = 150.0
)

//  360 * x     pi    6 * x     pi    pi * x
// -------- * ---- = ------ * ---- = -------
//       60    180             180        30

func SecondsInRad(tm time.Time) float64 {
	u := math.Pi * float64(tm.Second()) / 30
	return u
}

func MinutesInRad(tm time.Time) float64 {
	u := SecondsInRad(tm)/60 + (math.Pi*float64(tm.Minute()))/30
	return u
}

func HoursInRad(tm time.Time) float64 {
	u := MinutesInRad(tm)/12 + math.Pi*float64(tm.Hour()%12)/6
	return u
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func SecondHandPoint(tm time.Time) Point {
	angle := SecondsInRad(tm)
	return angleToPoint(angle)
}

func MinuteHandPoint(tm time.Time) Point {
	angle := MinutesInRad(tm)
	return angleToPoint(angle)
}

func HourHandPoint(tm time.Time) Point {
	angle := HoursInRad(tm)
	return angleToPoint(angle)
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}
	return p
}

func SecondHand(w io.Writer, tm time.Time) {
	p := makeHand(SecondHandPoint(tm), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func MinuteHand(w io.Writer, tm time.Time) {
	p := makeHand(MinuteHandPoint(tm), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func HourHand(w io.Writer, tm time.Time) {
	p := makeHand(HourHandPoint(tm), hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func SVGWriter(w io.Writer, tm time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, svgBezel)
	SecondHand(w, tm)
	MinuteHand(w, tm)
	HourHand(w, tm)
	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
const svgBezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgEnd = `</svg>`
