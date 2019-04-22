package main

import (
	"github.com/Myriad-Dreamin/go-rlp/rlp"
	"fmt"
	"io"
	"errors"
)

type Atom interface{}
type Mty struct {
	G Atom
}

func (mt *Mty) EncodeRLP(w io.Writer) (err error) {
	if mt == nil {
		return errors.New("nil node")
	}
	switch mt.G.(type) {
	
	case []Atom: // as list
		var lis = mt.G.([]Atom)
		var enc_lis = make([][]byte, 0)
		var enc_ele []byte
		for _, ele := range lis {
			enc_ele, err = rlp.EncodeToBytes(ele.(*Mty))
			if err != nil {
				return
			}
			enc_lis = append(enc_lis, enc_ele)
		}
		fmt.Println(enc_lis)
		enc_ele, err = rlp.EncodeToBytes(enc_lis)
		if err != nil {
			return
		}
		enc_lis = make([][]byte, 0)

		return rlp.Encode(w, append(enc_lis, enc_ele))
	
	case []byte:// as bytes
		return rlp.Encode(w, mt.G.([]byte))
		
	default:
		return errors.New("unrecognized Atom type")
	}
}

func main() {
	var mt *Mty
	bt, err := rlp.EncodeToBytes(mt)
	if err != nil {
		fmt.Println("Exception:", err)
	}
	fmt.Printf("%v -> %X \n", mt, bt)
	
	mt = &Mty{G:append(make([]Atom, 0), &Mty{G:[]byte("here")}, &Mty{G:[]byte("here")})}
	bt, err = rlp.EncodeToBytes(mt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v -> %X \n", mt, bt)

	var node Mty
	err = rlp.DecodeBytes(bt, &node)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("decode ret:", node, node.G)
	
	var dec_lis []Atom
	err = rlp.DecodeBytes(node.G.([]byte), &dec_lis)
	if err != nil {
		fmt.Println(err)
	}
	node.G = dec_lis
	fmt.Println("decode ret:", node, node.G)
}