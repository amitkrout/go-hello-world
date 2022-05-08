package pkg

import (
	"testing"
)

func TestAddtion(t *testing.T) {
	expect := Addition(5, 5)
	if expect != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", expect, 10)
	}
}
