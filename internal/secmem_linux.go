package internal

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func initEnvContents(size int) error {
	// TODO(lstoll) memfd_secret implementation for kernel > 5.14
	var err error
	EnvContents, err = unix.Mmap(-1, 0, int(size), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_PRIVATE|unix.MAP_ANONYMOUS)
	if err != nil {
		return fmt.Errorf("mmap: %w", err)
	}

	_ = unix.Madvise(EnvContents, unix.MADV_DONTDUMP)

	if err := unix.Mlock(EnvContents); err != nil {
		return fmt.Errorf("locking mem: $v")
	}

	return nil
}
