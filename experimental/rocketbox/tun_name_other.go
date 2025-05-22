//go:build !(darwin || linux)

package rocketbox

import "os"

func getTunnelName(fd int32) (string, error) {
	return "", os.ErrInvalid
}
