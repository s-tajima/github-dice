package main

import (
	"testing"
)

func TestInitialize(t *testing.T) {
	d := &Dice{}
	d.initialize([]string{})

	if d.Opts.Query != "is:issue" {
		t.Error("--query should be set 'is:issue'")
	}
}
