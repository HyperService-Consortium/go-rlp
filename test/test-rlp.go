package main

import rlp "./go-rlp"
import (
	"fmt"
	"io"
)

type Atom interface{}
type Mty struct {
	G Atom
	T uint
}

func (mt *Mty) EncodeRLP(w io.Writer) (err error) {
	if mt == nil {
		err = rlp.Encode(w, []uint{0})
	} else if mt.T == 1 {
		var bt []byte
		bt, err = rlp.EncodeToBytes(mt.G.(*Mty))
		if err != nil {
			return
		}
		err = rlp.Encode(w, bt)
	} else {
		err = rlp.Encode(w, mt.G)
		if err != nil {
			return
		}
	}
	return
}

type Mtx struct {
	Y []byte
	T int
}

func main() {
	var mt *Mty
	bt, err := rlp.EncodeToBytes(mt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v -> %X \n", mt, bt)
	fmt.Println(bt)
	
	mt = &Mty{G:&Mty{G:[]byte("here"), T:2}, T:1}
	bt, err = rlp.EncodeToBytes(mt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v -> %X \n", mt, bt)
	fmt.Println(bt)

	var mx * Mtx
	bt, err = rlp.EncodeToBytes(mx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v -> %X \n", mx, bt)
	mx = &Mtx{Y:[]byte("Here"), T:2}
	bt, err = rlp.EncodeToBytes(mx)
	fmt.Printf("%v -> %X \n", mx, bt)

}