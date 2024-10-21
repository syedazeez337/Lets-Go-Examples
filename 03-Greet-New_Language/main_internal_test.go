package main

import "testing"

func TestGreet_English(t *testing.T) {
	lang := language("en")
	want := "Hello world"

	got := Greet(lang)

	if got != want {
		t.Errorf("Expected: %q, Got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := language("fr")
	want := "Bonjour le monde"

	got := Greet(lang)

	if want != got {
		t.Errorf("Expected: %q, Got: %q", want, got)
	}
}

func TestGreet_Default(t *testing.T) {
	lang := language("akk")
	want := ""

	got := Greet(lang)

	if want != got {
		t.Errorf("Expected: %q, Got: %q", want, got)
	}
}
