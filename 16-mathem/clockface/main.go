package main

import (
	"mathem"
	"os"
	"time"
)

func main() {
	tm := time.Now()
	mathem.SVGWriter(os.Stdout, tm)
}
