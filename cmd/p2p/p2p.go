package p2p

import (
	"golang.org/x/net/websocket"

	"github.com/quantumshiro/perychain/pkg/blockchain"
)

type P2P struct {
	BlockChain *blockchain.BlockChain
	sockets    []*websocket.Conn
}

func NewP2P(bc *blockchain.BlockChain) *P2P {
	return &P2P{bc, []*websocket.Conn{}}
}

func listen(ws *websocket.Conn, p2p *P2P) {
	p2p.sockets = append(p2p.sockets, ws)
}

func connectPeers(p2p *P2P) {
	for _, ws := range p2p.sockets {
		for _, peer := range p2p.sockets {
			if ws == peer {
				continue
			}

			websocket.Message.Send(ws, "Hello, World!")
		}
	}
}

func ConnectSocket(ws *websocket.Conn, p2p *P2P) {
	listen(ws, p2p)
	connectPeers(p2p)
}
