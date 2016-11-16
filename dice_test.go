package main

import "testing"

func TestThrowWithoutExemptions(t *testing.T) {
	d := NewDice([]string{})
	c := []string{"foo"}
	if d.Throw(c) != "foo" {
		t.Error()
	}
}

func TestThrowWithExemptions(t *testing.T) {
	d := NewDice([]string{"foo"})
	c := []string{"foo", "bar"}
	for i := 0; i < 10; i++ {
		if d.Throw(c) == "foo" {
			t.Error()
		}
	}
}
