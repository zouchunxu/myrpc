package tests

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type demo struct {
	H string
}

func TestGobT(t *testing.T) {
	b := &bytes.Buffer{}
	w := gob.NewEncoder(b)
	r := gob.NewDecoder(b)

	var i int
	i = 10
	_ = w.Encode(&i)
	d := &demo{
		H: "hh",
	}
	_ = w.Encode(d)
	fmt.Println(string(b.Bytes()))
	var x int
	var dd demo
	_ = r.Decode(&x)

	_ = r.Decode(&dd)

	fmt.Println(x)
	fmt.Println(dd.H)
}
