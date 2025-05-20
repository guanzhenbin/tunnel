//go:build !unix

package rocketbox

import (
	"net"
)

func linkFlags(rawFlags uint32) net.Flags {
	panic("stub!")
}
