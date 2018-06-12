package main

import (
	"os"
	"testing"
)

var testCases = []struct {
	name string
	in map[string]string
	retVal int
} {
	{"help", map[string]string{"help": "true"}, 0},
	{"verbose", map[string]string{"verbose": "true"}, 0},
}

func TestRunServer(t *testing.T) {
	for _, tt := range testCases {
		for flag, flagVal := range tt.in {
			t.Logf("Setting %s to %s", flag, flagVal)
			os.Setenv(flag, flagVal)
		}

		t.Run(tt.name, func(t *testing.T) {
			retVal := runServer(true)
			if retVal != tt.retVal {
				t.Errorf("got %q, wanted %q", retVal, tt.retVal)
			}
		})

		for flag := range tt.in {
			t.Logf("Unsetting %s", flag)
			os.Unsetenv(flag)
		}
	}
}
