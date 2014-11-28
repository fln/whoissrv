package whoissrv

import (
	"net"
)

// The HandlerFunc type is an adapter to allow the use of ordinary functions as
// Handler interface implementations. If f is a function with the appropriate
// signature, HandlerFunc(f) is a Handler object that calls f.
type HandlerFunc func(string, net.Conn)

// HandleRequest calls f(req, c).
func (f HandlerFunc) HandleRequest(req string, c net.Conn) {
	f(req, c)
}
