package main

import (
	"os"
	"testing"
)

var testCases = []struct {
	name string
	in []string
	retVal int
} {
	{"help", []string{"--help", "true"}, 0},
	{"verbose", []string{"--verbose", "true"}, 0},
}

func TestRunServer(t *testing.T) {
	for _, tt := range testCases {
		flag := tt.in[0][2:]
		flagVal := tt.in[1]
		t.Logf("Setting %s to %s", flag, flagVal)
		os.Setenv(flag, flagVal)

		t.Run(tt.name, func(t *testing.T) {
			retVal := runServer(true)
			if retVal != tt.retVal {
				t.Errorf("got %q, wanted %q", retVal, tt.retVal)
			}
		})

		t.Logf("Unsetting %s", flag)
		os.Unsetenv(flag)
	}
}
