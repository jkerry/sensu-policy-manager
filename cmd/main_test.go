package sensupackagemanager

import (
	"testing"

	"rsc.io/quote"
)

func Hello() string {
	return quote.Hello()
}

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
