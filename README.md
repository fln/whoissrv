whoissrv
========

whoissrv is a golang library for writing whois (RFC 3912) servers. Design of
this library is similar to golang net/http package.

Library usage example:

```go
package main

import (
        "log"
        "fmt"
        "net"
        "github.com/fln/whoissrv"
)

func myFunc(req string, c net.Conn) {
        fmt.Fprintf(c, "Hello stranger from %v\r\n", c.RemoteAddr())
        fmt.Fprintf(c, "I know nothing about \"%v\"\r\n", req)
}

func main() {
        var srv whoissrv.Server

        srv.ListenAddr = ":4343"
        srv.Handler = whoissrv.HandlerFunc(myFunc)

        log.Panic(srv.ListenAndServe())
}

```

