package main

import (
	"fmt"
	"testing"
)

func count(s, r string) int {
	if len(s) == 0 {
		return 1
	}

	var res int
	for i := 0; i < len(r); i++ {
		if s[0] == r[i] {
			res += count(s[1:], r[i+1:])
		}
	}

	return res
}

func TestCount(t *testing.T) {
	var tt = []struct {
		s string
		r string
		e int
	}{
		{
			s: "bar",
			r: "barbar",
			e: 4,
		},
		{
			s: "abc",
			r: "abcdefab",
			e: 1,
		},
	}

	for _, test := range tt {
		c := count(test.s, test.r)
		if c != test.e {
			t.Fatalf("invalid count %d for %v", c, test)
		}
	}
}

func TestMain(t *testing.T) {
	fmt.Println("Hello Test")
}
