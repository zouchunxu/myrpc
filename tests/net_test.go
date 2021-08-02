package tests

import (
	"encoding/gob"
	"fmt"
	"net"
	"sync"
	"testing"
)

type Header struct {
	ServiceName string
	Args        interface{}
}

func TestConn(t *testing.T) {
	ch := make(chan struct{})
	//server

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		l, _ := net.Listen("tcp", ":9999")
		ch <- struct{}{}
		conn, _ := l.Accept()
		decoder := gob.NewDecoder(conn)
		sh := Header{}
		decoder.Decode(&sh)
		fmt.Printf("h: %+v\n", sh)
	}()
	<-ch
	//client
	dial, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	h := Header{
		ServiceName: "h.n",
		Args:        "aa",
	}
	encoder := gob.NewEncoder(dial)
	err = encoder.Encode(h)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wg.Wait()

}
