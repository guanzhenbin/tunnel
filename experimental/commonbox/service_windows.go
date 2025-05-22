//go:build windows

package commonbox

import "os"

func Dup(fd int) (nfd int, err error) {
	return 0, os.ErrInvalid
}
