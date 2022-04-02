package roman

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestRomanNums(t *testing.T) {

	testCases := []struct {
		InputNum  uint16
		OutputNum string
	}{
		{InputNum: 1, OutputNum: "I"},
		{InputNum: 2, OutputNum: "II"},
		{InputNum: 3, OutputNum: "III"},
		{InputNum: 4, OutputNum: "IV"},
		{InputNum: 5, OutputNum: "V"},
		{InputNum: 6, OutputNum: "VI"},
		{InputNum: 7, OutputNum: "VII"},
		{InputNum: 8, OutputNum: "VIII"},
		{InputNum: 9, OutputNum: "IX"},
		{InputNum: 10, OutputNum: "X"},
		{InputNum: 14, OutputNum: "XIV"},
		{InputNum: 18, OutputNum: "XVIII"},
		{InputNum: 20, OutputNum: "XX"},
		{InputNum: 39, OutputNum: "XXXIX"},
		{InputNum: 40, OutputNum: "XL"},
		{InputNum: 47, OutputNum: "XLVII"},
		{InputNum: 49, OutputNum: "XLIX"},
		{InputNum: 50, OutputNum: "L"},
		{InputNum: 100, OutputNum: "C"},
		{InputNum: 151, OutputNum: "CLI"},
		{InputNum: 500, OutputNum: "D"},
		{InputNum: 699, OutputNum: "DCXCIX"},
		{InputNum: 798, OutputNum: "DCCXCVIII"},
		{InputNum: 1000, OutputNum: "M"},
		{InputNum: 1984, OutputNum: "MCMLXXXIV"},
		{InputNum: 1986, OutputNum: "MCMLXXXVI"},
		{InputNum: 1999, OutputNum: "MCMXCIX"},
		{InputNum: 2021, OutputNum: "MMXXI"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("convert %d", testCase.InputNum), func(t *testing.T) {
			got, err := ConvertToRoman(testCase.InputNum)
			want := testCase.OutputNum
			if got != want {
				t.Errorf("expected %q got %q", want, got)
			}
			if err != nil {
				t.Errorf("no error expected but thrown")
			}
		})
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("convert %s", testCase.OutputNum), func(t *testing.T) {
			got := ConvertToArabic(testCase.OutputNum)
			want := testCase.InputNum
			if got != want {
				t.Errorf("expected %d got %d", want, got)
			}
		})
	}

	t.Run("convert 4000", func(t *testing.T) {
		_, err := ConvertToRoman(4000)
		if err == nil {
			t.Error("expected error but none thrown")
		}
	})
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			//log.Println(arabic)
			return true
		}
		roman, _ := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed check", err)
	}
}
