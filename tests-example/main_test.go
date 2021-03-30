package main

import (
	"testing"
	"time"
)

func TestUpdateTimeDisplayNonzero(t *testing.T) {
	r, err := NewExampleRouter()
	if err != nil {
		t.Fatalf("error making new router: %v", err)
	}

	r.updateDuration = time.Second
	actual := r.updateTimeDisplay()
	expected := "1s"

	if expected != actual {
		t.Fatalf("expected update duration %s but got %s", expected, actual)
	}
}

func TestUpdateTimeNonzero(t *testing.T) {
	r, err := NewExampleRouter()
	if err != nil {
		t.Fatalf("error making new router: %v", err)
	}

	r.updateDuration = 0
	actual := r.updateTimeDisplay()
	expected := "N/A"

	if expected != actual {
		t.Fatalf("expected update duration %s but got %s", expected, actual)
	}
}

func TestUpdateTimeTruncate(t *testing.T) {
	r, err := NewExampleRouter()
	if err != nil {
		t.Fatalf("error making new router: %v", err)
	}

	r.updateDuration = time.Millisecond * 123
	actual := r.updateTimeDisplay()
	expected := "100ms"  // truncate to nearest 100ms

	if expected != actual {
		t.Fatalf("expected update duration %s but got %s", expected, actual)
	}
}
