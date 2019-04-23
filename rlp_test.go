package rlp

import (
	_ "fmt"
	"io"
	"errors"
	"testing"
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
			enc_ele, err = EncodeToBytes(ele.(*Mty))
			if err != nil {
				return
			}
			enc_lis = append(enc_lis, enc_ele)
		}
		// fmt.Println(enc_lis)
		enc_ele, err = EncodeToBytes(enc_lis)
		if err != nil {
			return
		}
		enc_lis = make([][]byte, 0)

		return Encode(w, append(enc_lis, enc_ele))
	
	case []byte:// as bytes
		return Encode(w, mt.G.([]byte))
		
	default:
		return errors.New("unrecognized Atom type")
	}
}

func TestRLP(t *testing.T) {
	var mt = &Mty{G:append(make([]Atom, 0), &Mty{G:[]byte("here")}, &Mty{G:[]byte("here")})}
	bt, err := EncodeToBytes(mt)
	if err != nil {
		t.Error(err)
		return 
	}
	// fmt.Printf("%v -> %X \n", mt, bt)

	var node Mty
	err = DecodeBytes(bt, &node)
	if err != nil {
		t.Error(err)
		return 
	}

	// fmt.Println("decode ret:", node, node.G)
	
	var dec_lis []Atom
	err = DecodeBytes(node.G.([]byte), &dec_lis)
	if err != nil {
		t.Error(err)
		return 
	}
	node.G = dec_lis
	// fmt.Println("decode ret:", node, node.G)
}