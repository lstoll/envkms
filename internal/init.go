package internal

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

const EnvVar = "ENVKMS_KEYSET"

var EnvContents []byte

func init() {
	// we read and unset the env var as soon as possible, to minimise it's
	// access from anywhere else. We don't process it though, because we want
	// the user to be able to handle errors etc.
	ks, ok := os.LookupEnv(EnvVar)
	if !ok {
		return
	}
	if err := initEnvContents(len(ks)); err != nil {
		panic(fmt.Sprintf("allocating memory for env var: %v", err))
	}
	if _, err := io.ReadFull(strings.NewReader(ks), EnvContents); err != nil {
		panic(fmt.Sprintf("writing env contents to buf: %v", err))
	}
	if err := os.Unsetenv(EnvVar); err != nil {
		panic(fmt.Sprintf("unsetting env var: %v", err))
	}
	runtime.GC()
}
