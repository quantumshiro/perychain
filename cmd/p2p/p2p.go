package p2p

import (
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/quantumshiro/perychain/pkg/blockchain"
)

type P2P struct {
	BlockChain *blockchain.BlockChain
	peers      []*websocket.Conn
}

func NewP2P(bc *blockchain.BlockChain) *P2P {
	return &P2P{bc, []*websocket.Conn{}}
}

func (p2p *P2P) AddPeer(ws *websocket.Conn) {
	p2p.peers = append(p2p.peers, ws)
}

func (p2p *P2P) Broadcast() {
	for _, peer := range p2p.peers {
		peer.Write([]byte("Hello, World!"))
	}
}

func (p2p *P2P) Handle(ws *websocket.Conn) {
	p2p.AddPeer(ws)
	p2p.Broadcast()
}

func (p2p *P2P) Start() {
	http.Handle("/p2p", websocket.Handler(p2p.Handle))
	http.ListenAndServe(":8080", nil)
}
