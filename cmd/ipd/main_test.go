package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunServerHelp(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	log.Println("Testing server with `--help` flag")
	os.Args = []string{"noop", "--help"}
	retVal := runServer()

	assert.Equal(t, 0, retVal, "Expect a return value of zero.")
}