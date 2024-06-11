package test

import (
	"reflect"
	"testing"

	"github.com/quantumshiro/perychain/pkg/blockchain"
)

func TestAddBlock(t *testing.T) {
	bc := blockchain.NewBlockChain()
	bc.AddBlock([]byte("foo"))

	if !reflect.DeepEqual(bc.Chain[1].Data, []byte("foo")) {
		t.Errorf("BlockChain.AddBlock().Data = %s, want %s", bc.Chain[1].Data, []byte("foo"))
	}
}
