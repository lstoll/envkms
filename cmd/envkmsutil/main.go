package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/lstoll/envkms/internal"
)

func main() {
	if len(os.Args) < 2 {
		fatalf("subcommand required")
	}

	switch os.Args[1] {
	default:
		fatalf("%s is an invalid subcommand", os.Args[1])
	case "generate-keyset":
		s, err := internal.GenerateKeyset()
		if err != nil {
			fatalf("generating keyset: %v", err)
		}
		fmt.Printf("keyset: %s\n", s)
	}
}

func fatalf(format string, a ...any) {
	slog.Error(fmt.Sprintf(format, a...))
	os.Exit(1)
}
