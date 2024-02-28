package internal

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/unix"
)

func allocProtectedMem(size int) ([]byte, error) {
	// TODO(lstoll) memfd_secret implementation for kernel > 5.14
	data, err := unix.Mmap(-1, 0, int(size), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_PRIVATE|unix.MAP_ANONYMOUS)
	if err != nil {
		return nil, fmt.Errorf("mmap: %w", err)
	}

	runtime.SetFinalizer(&data, func(b *[]byte) {
		for i := range *b {
			(*b)[i] = byte(0)
		}
	})

	return data, nil
}
