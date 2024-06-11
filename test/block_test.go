package test

import (
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
	want := block.NewBlock(1, []byte{}, []byte{}, []byte{}).String()

	if b.String() != want {
		t.Errorf("Block.Genesis() = %s, want %s", b.String(), want)
	}
}
