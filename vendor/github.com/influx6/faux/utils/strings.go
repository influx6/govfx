package utils

import "strconv"

// StringMatcher defines a function type to match a giving string against.
type StringMatcher func(string) bool

// MatchAny checks if the giving string matches any of the possiblities, it
// converts the giving type as necessary and if the possiblity is a function
// runs the function with the expected signature of `MatchString`.
// If runs until it finds a match then stops else returns false if non matches.
// Only checks for int,uint,float,rune,string and StringMatcher types.
func MatchAny(target string, possibilities ...interface{}) bool {
	var state bool

	for _, item := range possibilities {
		if state {
			break
		}

		switch item.(type) {
		case StringMatcher:
			if (item.(StringMatcher))(target) {
				state = true
				continue
			}
		case rune:
			if string(item.(rune)) == target {
				state = true
				continue
			}
		case string:
			if item.(string) == target {
				state = true
				continue
			}
		case int, int64:
			if target == strconv.FormatInt(item.(int64), 10) {
				state = true
				continue
			}
		case uint, uint32, uint64:
			if target == strconv.FormatUint(item.(uint64), 10) {
				state = true
				continue
			}
		case float32, float64:
			if target == strconv.FormatFloat(item.(float64), 'f', 1, 64) {
				state = true
				continue
			}
		}
	}

	return state
}

// MatchAll checks if the giving string matches all the provided possibilities
// else returns false. It converts the possiblity into a string when possible
// and if its a StringMatcher type then runs the giving function.
// Only checks for int,uint,float,rune,string and StringMatcher types.
func MatchAll(target string, possibilities ...interface{}) bool {
	state := true

	for _, item := range possibilities {
		if !state {
			break
		}

		switch item.(type) {
		case StringMatcher:
			if !(item.(StringMatcher))(target) {
				state = false
				continue
			}
		case rune:
			if string(item.(rune)) != target {
				state = false
				continue
			}
		case string:
			if item.(string) != target {
				state = false
				continue
			}
		case int, int64:
			if target != strconv.FormatInt(item.(int64), 10) {
				state = false
				continue
			}
		case uint, uint32, uint64:
			if target != strconv.FormatUint(item.(uint64), 10) {
				state = false
				continue
			}
		case float32, float64:
			if target == strconv.FormatFloat(item.(float64), 'f', 1, 64) {
				state = false
				continue
			}
		}
	}

	return state
}
