package go_say_hello

import "testing"

func TestSayHello(t *testing.T) {
	result := SayHello("World")
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("SayHello('World') = %v; want %v", result, expected)
	}
}

func TestSayHelloEmpty(t *testing.T) {
	result := SayHello("")
	expected := "Hello, !"
	if result != expected {
		t.Errorf("SayHello('') = %v; want %v", result, expected)
	}
}

func TestSayHelloSpecialChars(t *testing.T) {
	result := SayHello("Gopher!")
	expected := "Hello, Gopher!!"
	if result != expected {
		t.Errorf("SayHello('Gopher!') = %v; want %v", result, expected)
	}
}