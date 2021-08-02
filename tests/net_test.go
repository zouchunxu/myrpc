package tests

import (
	"encoding/gob"
	"net"
	"testing"
)

type header struct {
	serviceName string
	args        interface{}
}

func TestHello(t *testing.T) {

	//server
	go func() {
		l, _ := net.Listen("tcp", ":9999")
		conn, _ := l.Accept()
		decoder := gob.NewDecoder(conn)
		h := header{}
		decoder.Decode(&h)
	}()

	//client

}
