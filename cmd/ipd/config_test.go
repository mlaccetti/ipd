package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	opts := config()
	assert.NotNil(t, opts, "opt should not be nil")
	fmt.Printf("%+v\n", opts)
}

func TestOutput(t *testing.T) {
	printHelp()
}