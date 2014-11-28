package whoissrv

import (
	"net"
	"bufio"
	"time"
	"strings"
)

// Server defines parameters for running a whois server. The zero value for
// Server is a valid configuration.
type Server struct {
	ListenAddr string         // TCP address to listen on, ":43" if empty
	ReadTimeout time.Duration // Maximum duration client have to send request before timing out
	Handler Handler           // handler to invoke when request ir parsed
}

// ListenAndServe listens on the TCP network address this.ListenAddr and then 
// calls Serve to handle requests on incoming connections. If this.ListenAddr
// is empty, ":43" is used.
func (this *Server) ListenAndServe() error {
	if this.ListenAddr == "" {
		this.ListenAddr = ":43"
	}

	l, err := net.Listen("tcp", this.ListenAddr)
	if err != nil {
		return err
	}
	return this.Serve(l)
}

// Serve accepts incoming connections on the Listener l, creating a new service
// goroutine for each. The service goroutines read requests and then call
// this.Handler to reply to them. 
func (this *Server) Serve(l net.Listener) error {
	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}
		go this.processRequest(c)
	}
}

func (this *Server) processRequest(c net.Conn) {
	var br = bufio.NewReader(c)

	defer c.Close()

	if this.ReadTimeout != 0 {
		c.SetDeadline(time.Now().Add(this.ReadTimeout))
	}

	line, err := br.ReadString('\n')
	if err != nil {
		return
	}

	if this.Handler != nil {
		this.Handler.HandleRequest(strings.TrimRight(line, "\r\n"), c)
	}
}

