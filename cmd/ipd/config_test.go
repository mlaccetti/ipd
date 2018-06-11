package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	viper, opts, err := config()
	assert.Nil(t, err, "loading a config should not throw an err")
	assert.NotNil(t, viper, "viper should not be nil")
	assert.NotNil(t, opts, "opt should not be nil")
	fmt.Printf("%+v\n", opts)
}

func TestOutput(t *testing.T) {
	printHelp()
}