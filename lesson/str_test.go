package lesson

import "testing"

func TestReverse(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Backward", "drawkcaB"},
		{"hello, 小妞", "妞小 ,olleh"},
		{"", ""},
	}
	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q)== %q,want %q", c.s, got, c.want)
		}
	}
}
