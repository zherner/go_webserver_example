package main

import "testing"

//
func Test_config(t *testing.T) {
	want := cfg{address: "", port: ""}

	got := config()

	if got != want {
		t.Errorf("Test_config(%q) == %q, want %q", want, got, want)
		return
	}
}
