package grep

import (
	"DistributedGrep/internal/app"
	"testing"
)

func TestMatch(t *testing.T) {

	req := app.GrepRequest{
		Pattern: "ERROR",
	}

	if !Match("ERROR database", req) {
		t.Fatal("expected match")
	}

	if Match("INFO started", req) {
		t.Fatal("unexpected match")
	}
}

func TestIgnoreCase(t *testing.T) {

	req := app.GrepRequest{
		Pattern:    "error",
		IgnoreCase: true,
	}

	if !Match("ERROR database", req) {
		t.Fatal("ignore case failed")
	}
}

func TestInvert(t *testing.T) {

	req := app.GrepRequest{
		Pattern: "ERROR",
		Invert:  true,
	}

	if Match("ERROR database", req) {
		t.Fatal("invert failed")
	}

	if !Match("INFO started", req) {
		t.Fatal("invert failed")
	}
}