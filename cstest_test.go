package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("Hello Test")
}

func TestFailure(t *testing.T) {
	t.Fail()
}
