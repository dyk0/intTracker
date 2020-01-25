package stringparse

import "testing"

func TestParse(t *testing.T) {
	cases := []struct {
		in string; want bool
	}{
		{"foo_bar-baz", true},
		{"foo_bar baz", false},
		{"foo bar-baz", false},
		{"foo-bar_baz1", true},
	}
	for _, c := range cases {
		got := Parse(c.in)
		if got != c.want {
			t.Errorf("Parse(%v) == %v, want %v", c.in, got, c.want)
		}
	}

}
