package internal

import (
	"runtime"
	"testing"
)

func TestSecMem(t *testing.T) {
	// quick smoke test to make sure it doesn't panic.
	if err := initEnvContents(64); err != nil {
		t.Fatal(err)
	}
	runtime.GC()
}
