package tt

import (
	"testing"
)

// AssertEqual validates that val1 is equal to val2 and throws an error with line number
func AssertEqual(t *testing.T, expected, actual interface{}) {
	equalSkip(t, 2, expected, actual)
}

// AssertNotEqual validates that val1 is not equal val2 and throws an error with line number
func AssertNotEqual(t *testing.T, unexpected, actual interface{}) {
	notEqualSkip(t, 2, unexpected, actual)
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

func AssertIsNil(t *testing.T, actual interface{}) {
	equalSkip(t, 2, nil, actual)
}

func AssertIsNotNil(t *testing.T, actual interface{}) {
	notEqualSkip(t, 2, nil, actual)
}

func AssertTrue(t *testing.T, actual bool) {
	equalSkip(t, 2, true, actual)
}

func AssertFalse(t *testing.T, actual bool) {
	equalSkip(t, 2, false, actual)
}

func AssertInMap(t *testing.T, m interface{}, key interface{}) {
	inMapSkip(t, 2, m, key)
}
