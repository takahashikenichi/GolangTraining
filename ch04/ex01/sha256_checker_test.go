package src

import (
	"testing"
)

func TestSha256Compare(t *testing.T) {
	t.Parallel()
	diff := Sha256Compare([32]byte{byte(255)}, [32]byte{byte(0)})
	if diff != 8{
		t.Errorf("Difference is not 8 , %d", diff)
	}
}

func TestSha256Checker(t *testing.T) {
	t.Parallel()
	diff := Sha256Checker("x","X")
	if diff != 125{
		t.Errorf("Difference is not 125 , %d", diff)
	}
}

