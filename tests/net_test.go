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
	Args        Zx
}

type Zcx struct {
}

type Zx struct {
	A string
}

func (Zcx) SayDemo(s Zx) int {
	fmt.Println(s.A)
	return 2
}

func TestConn(t *testing.T) {
	ch := make(chan struct{})
	//server
	mp := sync.Map{}
	z := &Zcx{}
	mp.Store(reflect.ValueOf(z).Type().String(), z)

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

		var in []reflect.Value
		in = append(in, reflect.ValueOf(sh.Args))
		z, ok := mp.Load(sh.ServiceName)
		if !ok {
			return
		}
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
	tz := Zx{
		A: "zzz",
	}
	h := Header{
		ServiceName: reflect.ValueOf(z).Type().String(),
		Args:        tz,
	}
	encoder := gob.NewEncoder(dial)
	err = encoder.Encode(h)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wg.Wait()
}
