package main

import "testing"

func TestWaitForServer(t *testing.T) {
	tests := []struct {
		url  string
		want bool
	}{
		{"http://baddk.gopl.org", false},
		{"http://gopl.org", true},
	}

	for _, test := range tests {
		if got := WaitForServer(test.url); (got == nil) != test.want {
			t.Errorf("WaitForServer(%q) = %v", test.url, got)
		}
	}
}
