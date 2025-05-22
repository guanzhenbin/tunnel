//go:build !windows

package commonbox

import "syscall"

func Dup(fd int) (nfd int, err error) {
	return syscall.Dup(fd)
}
