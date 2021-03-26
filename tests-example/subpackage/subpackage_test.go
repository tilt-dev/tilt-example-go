package subpackage

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	fmt.Println("foo is working")
}

func TestBar(t *testing.T) {
	fmt.Println("bar is working")
}

func TestBaz(t *testing.T) {
	fmt.Println("baz is working")
}

func TestPicard(t *testing.T) {
	numLights := 4
	if numLights != 5 {
		t.Fatal("there are five lights")
	}
}
