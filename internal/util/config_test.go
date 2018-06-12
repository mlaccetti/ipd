package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	viper := Config()
	assert.NotNil(t, viper, "viper should not be nil")
}

func TestOutput(t *testing.T) {
	PrintHelp()
}