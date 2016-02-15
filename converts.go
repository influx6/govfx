package vfx

import (
	"regexp"
	"strconv"
)

//==============================================================================

// nodigits defines a regexp for matching non-digits.
var nodigits = regexp.MustCompile("[^\\d]+")

// ParseFloat parses a string into a float if fails returns the default value 0.
func ParseFloat(fl string) float64 {
	fll, _ := strconv.ParseFloat(DigitsOnly(fl), 64)
	return fll
}

// ParseInt parses a string into a int if fails returns the default value 0.
func ParseInt(fl string) int {
	fll, _ := strconv.Atoi(DigitsOnly(fl))
	return fll
}

// ParseIntBase16 parses a string into a int using base16.
func parseIntBase16(fl string) int {
	fll, _ := strconv.ParseInt(fl, 16, 64)
	return int(fll)
}

// DigitsOnly removes all non-digits characters in a string.
func DigitsOnly(fl string) string {
	return nodigits.ReplaceAllString(fl, "")
}

//==============================================================================
