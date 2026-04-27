package golden

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func AssertEqualFile(t *testing.T, expectedPath string, actual string) {
	t.Helper()
	data, err := os.ReadFile(expectedPath)
	if err != nil {
		t.Fatalf("read golden file: %v", err)
	}
	expected := normalize(string(data))
	actual = normalize(actual)
	if expected == actual {
		return
	}
	t.Fatalf("golden mismatch for %s\n%s", expectedPath, firstDiff(expected, actual))
}

func normalize(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	s = strings.TrimRight(s, "\n") + "\n"
	return s
}

func firstDiff(expected, actual string) string {
	expLines := strings.Split(expected, "\n")
	actLines := strings.Split(actual, "\n")
	limit := len(expLines)
	if len(actLines) > limit {
		limit = len(actLines)
	}
	for i := 0; i < limit; i++ {
		var exp, act string
		if i < len(expLines) {
			exp = expLines[i]
		}
		if i < len(actLines) {
			act = actLines[i]
		}
		if exp != act {
			return fmt.Sprintf("first difference at line %d\nexpected: %q\nactual:   %q", i+1, exp, act)
		}
	}
	return "outputs differ"
}
