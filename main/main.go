package main

import (
	"context"
	"fmt"
	"myrpc"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func (f Foo) Sleep(args Args, reply *int) error {
	time.Sleep(time.Second * time.Duration(args.Num1))
	*reply = args.Num1 + args.Num2
	return nil
}

func (f Foo) Say(args string, reply *int) error {
	reply = new(int)
	t, _ := strconv.Atoi(args)
	reply = &t
	return nil
}

func main() {
	go func() {
		var f Foo
		l, _ := net.Listen("tcp", "127.0.0.1:9292")
		myrpc.Register(&f)
		myrpc.Accept(l)
	}()
	time.Sleep(time.Second * 2)
	c, _ := myrpc.Dial("tcp", "127.0.0.1:9292")

	var ret int
	args := &Args{
		Num1: 1,
		Num2: 1,
	}
	_ = c.Call(context.Background(), "Foo.Sum", args, &ret)
	fmt.Println(ret)

	l, _ := net.Listen("tcp", "127.0.0.1:9293")
	myrpc.HandleHTTP()

	http.Serve(l, nil)
}
