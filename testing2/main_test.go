package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(42, 16)
	if total != 58 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 58)
	}
}

func TestSubtract(t *testing.T) {
	total := Subtract(42, 16)
	if total != 26 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 26)
	}
}

func TestDoMath(t *testing.T) {
	x := DoMath(42, 16, Add)
	if x != 58 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", x, 58)
	}
	y := DoMath(42, 16, Subtract)
	if y != 26 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", y, 26)
	}
}

func TestParadise(t *testing.T) {
	got := paradise("Hawaii")
	want := "My idea of paradise is Hawaii"
	if got != want {
		log.Fatalf("error - want: %s, got: %s.", want, got)
	}
}

/*func TestAdd(t *testing.T) {
	got := add(7, 5)
	want := 12
	if got != want {
		log.Fatalf("error - want %v and got %v", want, got)
	}
}*/
