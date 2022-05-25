package Utils

import (
		"testing"
		)

func TestAtoi(t *testing.T) {
	var tests = []struct {
        input string
        want int
    }{
		{"123",123},
		{"12",12},
	}

	for _, test := range tests {
		if got:=Atoi(test.input) ; got != test.want {
			t.Errorf("InArray(%s) = %v  Not  %v", test.input , got , test.want )
		}
	}
}

