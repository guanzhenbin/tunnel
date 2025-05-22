//go:build !unix

package commonbox

import (
	"net"
)

func LinkFlags(rawFlags uint32) net.Flags {
	panic("stub!")
}
