package main

import "testing"

func TestGreet_English(t *testing.T) {
	// setup phase
	lang := language("en")
	want := "Hello world"

	// execution phase
	got := greet(lang)

	//validation phase
	if got != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le monde"

	got := greet(lang)

	if got != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	// Akkadian is not implemented yet!
	lang := language("akk")
	want := ""

	// execution phase
	got := greet(lang)

	//validation phase
	if got != want {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
