package main

import "testing"

func TestProcessInput(t *testing.T) {
	cases := []struct {
		in   string
		want float64
	}{
		{"R2, L3", 5.0},
		{"R2, R2, R2", 2.0},
		{"R5, L5, R5, R3", 12.0},
		{"R5, R5, R5, R5", 0.0},
		{"L5, L5, L5, L5", 0.0},
		{"L50", 50.0},
	}
	for _, c := range cases {
		got := ProcessInput(c.in)
		if got != c.want {
			t.Errorf("ProcessInput(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
