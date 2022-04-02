package iterations

import "strings"

func Repeat(what string, repeatFactor int) string {
	var repeated string
	for i := 0; i < repeatFactor; i++ {
		repeated += what
	}
	return repeated
}

func RepeatStd(what string, repeatFactor int) string {
	return strings.Repeat(what, repeatFactor)
}

func main() {

}
