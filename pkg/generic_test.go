package pkg

import (
	"testing"
)

func TestAddtion(t *testing.T) {
	got := Addition(5, 5)
	if got != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, 10)
	}
}
