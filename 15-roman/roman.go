package roman

import (
	"errors"
	"strings"
)

type RomanNum struct {
	Value  uint16
	Symbol string
}

type RomanNums []RomanNum

var allRomanNumbers = RomanNums{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(n uint16) (string, error) {
	var result strings.Builder

	if n > 3999 {
		return "", errors.New("unable to convert")
	}

	for _, num := range allRomanNumbers {
		for n >= num.Value {
			result.WriteString(num.Symbol)
			n -= num.Value
		}
	}

	return result.String(), nil
}

func (r RomanNums) ValueOf(n ...byte) uint16 {
	sym := string(n)
	for _, s := range r {
		if s.Symbol == sym {
			return s.Value
		}
	}
	return 0
}

func isSubstractor(sym uint8) bool {
	return sym == 'I' || sym == 'X' || sym == 'C' || sym == 'M'
}

func ConvertToArabic(n string) uint16 {
	var total uint16 = 0

	for i := 0; i < len(n); i++ {
		sym := n[i]
		notEnd := i+1 < len(n)

		if notEnd && isSubstractor(sym) {
			if val := allRomanNumbers.ValueOf(sym, n[i+1]); val != 0 {
				total += val
				i++
				continue
			}
		}
		total += allRomanNumbers.ValueOf(sym)
	}
	return total
}
