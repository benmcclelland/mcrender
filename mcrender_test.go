package mcrender

import (
	"os"
	"testing"
)

func TestCreateSTLFromInput(t *testing.T) {
	fname := "test.mcfunction"
	f, err := os.Open(fname)
	if err != nil {
		t.Errorf("open %v: %v", fname, err)
	}
	err = CreateSTLFromInput(f, "test.stl")
	if err != nil {
		t.Errorf("create STL: %v", err)
	}
}
