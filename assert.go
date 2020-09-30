package tt

import (
	"testing"
)

// AssertEqual validates that val1 is equal to val2 and throws an error with line number
func AssertEqual(t *testing.T, val1, val2 interface{}) {
	equalSkip(t, 2, val1, val2)
}

// AssertNotEqual validates that val1 is not equal val2 and throws an error with line number
func AssertNotEqual(t *testing.T, val1, val2 interface{}) {
	notEqualSkip(t, 2, val1, val2)
}

// AssertMatchRegex validates that value matches the regex, either string or *regex
// and throws an error with line number
func AssertMatchRegex(t *testing.T, value string, regex interface{}) {
	matchRegexSkip(t, 2, value, regex)
}

// AssertNotMatchRegex validates that value matches the regex, either string or *regex
// and throws an error with line number
func AssertNotMatchRegex(t *testing.T, value string, regex interface{}) {
	notMatchRegexSkip(t, 2, value, regex)
}

// AssertPanicMatches validates that the panic output of running fn matches the supplied string
func AssertPanicMatches(t *testing.T, fn func(), matches string) {
	panicMatchesSkip(t, 2, fn, matches)
}
