package tests

import (
	"encoding/gob"
	"fmt"
	"net"
	"reflect"
	"sync"
	"testing"
)

type Header struct {
	ServiceName string
	Args        interface{}
}

type Zcx struct {
}

func (Zcx) SayDemo(s string) int {
	fmt.Println(s)
	return 2
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

		z := &Zcx{}

		var in []reflect.Value
		in = append(in, reflect.ValueOf(sh.Args))
		reflect.ValueOf(z).MethodByName("SayDemo").Call(in)
		fmt.Println(reflect.ValueOf(z).MethodByName("SayDemo").Type().In(0).Name())
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
