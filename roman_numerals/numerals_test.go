package main

import (
	"log"
	"testing"
	"testing/quick"
)

var (
	cases = []struct {
		arabic uint16
		roman  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{14, "XIV"},
		{15, "XV"},
		{16, "XVI"},
		{25, "XXV"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{798, "DCCXCVIII"},
		{900, "CM"},
		{1000, "M"},
		{1006, "MVI"},
		{1984, "MCMLXXXIV"},
		{2021, "MMXXI"},
		{3999, "MMMCMXCIX"},
	}
)

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		got := ConvertToRoman(test.arabic)
		if got != test.roman {
			t.Errorf("got %q, want %q", got, test.roman)
		}
	}
}

func TestArabic(t *testing.T) {
	for _, test := range cases {
		got := ConvertToArabic(test.roman)
		if got != test.arabic {
			t.Errorf("got %d, want %d", got, test.arabic)
		}
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
