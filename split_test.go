package split

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSplit(t *testing.T) {
	test := func(name, input, sep string, want []string) {
		t.Run(name, func(t *testing.T) {
			got := Split(input, sep)
			diff := cmp.Diff(want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}

	test("simple", "a/b/c", "/", []string{"a", "b", "c"})
	test("wrong sep", "a/b/c", ",", []string{"a/b/c"})
	test("no sep", "abc", "/", []string{"abc"})
	test("trailing sep", "a/b/c/", "/", []string{"a", "b", "c", ""})
}
