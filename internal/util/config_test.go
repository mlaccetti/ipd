package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	viper, err := config()
	assert.Nil(t, err, "loading a util should not throw an err")
	assert.NotNil(t, viper, "viper should not be nil")
}

func TestOutput(t *testing.T) {
	printHelp()
}