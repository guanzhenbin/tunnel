//go:build !windows

package rocketbox

import "syscall"

func dup(fd int) (nfd int, err error) {
	return syscall.Dup(fd)
}
