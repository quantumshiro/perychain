package test

import (
	"reflect"
	"testing"

	"github.com/quantumshiro/perychain/pkg/block"
)

func TestBlock(t *testing.T) {
	b := block.NewBlock(1, []byte("foo"), []byte("bar"), []byte("baz"))
	want := block.NewBlock(1, []byte("foo"), []byte("bar"), []byte("baz")).String()

	if b.String() != want {
		t.Errorf("Block.String() = %s, want %s", b.String(), want)
	}
}

func TestGenesis(t *testing.T) {
	b := block.Genesis()

	if b.String() != block.NewBlock(1, []byte{}, []byte{}, []byte{}).String() {
		t.Errorf("Block.Genesis() = %s, want %s", b.String(), block.NewBlock(1, []byte{}, []byte{}, []byte{}).String())
	}
}

func TestMineBlock(t *testing.T) {
	b := block.NewBlock(1, []byte("foo"), []byte("bar"), []byte("baz"))
	newBlock := block.MineBlock(b, []byte("qux"))

	if !reflect.DeepEqual(newBlock.LastHash, b.Hash) {
		t.Errorf("MineBlock().LastHash = %s, want %s", newBlock.LastHash, b.Hash)
	}

	if !reflect.DeepEqual(newBlock.Data, []byte("qux")) {
		t.Errorf("MineBlock().Data = %s, want %s", newBlock.Data, []byte("qux"))
	}
}
