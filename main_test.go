package main

import "testing"

func Test_sayHello(t *testing.T) {
	name := "Rajesh"
	want := "Hi Rajesh"

	if got := sayHello(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}
