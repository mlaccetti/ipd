package main

import (
	"os"
	"testing"
)

var testCases = []struct {
	name   string
	in     map[string]string
	retVal int
}{
	{"help", map[string]string{"help": "true"}, 0},
	{"verbose", map[string]string{"verbose": "true"}, 0},
	{"country-db", map[string]string{"verbose": "true", "country-db": "../../data/country.mmdb"}, 0},
	{"city-db", map[string]string{"verbose": "true", "city-db": "../../data/city.mmdb"}, 0},
	{"country-and-city-db", map[string]string{"verbose": "true", "country-db": "../../data/country.mmdb", "city-db": "../../data/city.mmdb"}, 0},
	{"tls-enabled-but-broken", map[string]string{"verbose": "true", "listen-tls": ":8443"}, 0},
	{"tls-enabled-missing-cert-key", map[string]string{"verbose": "true", "listen-tls": ":8443", "tls-cert": "", "tls-key": ""}, 0},
	{"tls-enabled", map[string]string{"verbose": "true", "listen-tls": ":8443", "tls-cert": "../../certs/test-localhost.crt", "tls-key": "../../certs/test-localhost.key"}, 0},
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
				t.Errorf("got %d, wanted %d", retVal, tt.retVal)
			}
		})

		for flag := range tt.in {
			t.Logf("Unsetting %s", flag)
			os.Unsetenv(flag)
		}
	}
}
