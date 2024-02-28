package internal

import (
	"runtime"
	"testing"
)

func TestSecMem(t *testing.T) {
	// quick smoke test to make sure it doesn't panic.
	buf, err := allocProtectedMem(64)
	if err != nil {
		t.Fatal(err)
	}
	_ = buf
	buf = nil
	_ = buf
	runtime.GC()
}
