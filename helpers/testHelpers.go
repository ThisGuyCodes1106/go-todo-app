package testHelpers

import (
	"testing"
)

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
func AssertError(t testing.TB, err, want error) {
	t.Helper()
	if err != want {
		t.Errorf("got %s, want %s", err, want)
	}
}
func AssertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}