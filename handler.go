package whoissrv

import (
	"net"
)

type Handler interface {
	// HandleRequest is called when whois request arrives. Request string is
	// stored in "req" variable. Argument "c" is connection struct.
	HandleRequest(req string, c net.Conn)
}
