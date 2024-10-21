package main

import "testing"

/*
func ExampleGreet() {
	main()
	greet()
	// Output: Hello Greetings
}
*/

func TestGreet(t *testing.T) {
	want := "Hello Greetings"

	got := greet()

	if got != want {
		t.Errorf("Want: %q, got: %q", want, got)
	}
}
