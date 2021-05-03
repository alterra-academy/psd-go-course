package calculate

import (
	"testing"
)

func TestAdditon(t *testing.T) {
	if Additon(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Additon(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}
