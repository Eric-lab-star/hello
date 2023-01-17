package hello

import "testing"

func TestHello(t *testing.T) {
	in, want := "Eric", "Hello, Eric"
	if got := Hello(in); got != want {
		t.Errorf(`Hello(%s) want %s got %s`, in, want, got)
	}
}
