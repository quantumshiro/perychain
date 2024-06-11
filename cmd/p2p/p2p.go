package p2p

import (
	"net/http"

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

func (p2p *P2P) Blocks(ws *websocket.Conn) {
	p2p.sockets = append(p2p.sockets, ws)

	// send the chain to the new peer
	websocket.JSON.Send(ws, p2p.BlockChain.Chain)
}

func (p2p *P2P) Mine(data []byte) {
	p2p.BlockChain.AddBlock(data)

	// broadcast the new chain to all peers
	for _, ws := range p2p.sockets {
		websocket.JSON.Send(ws, p2p.BlockChain.Chain)
	}
}

func (p2p *P2P) Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (p2p *P2P) Start() {
	http.Handle("/blocks", websocket.Handler(p2p.Blocks))
	http.HandleFunc("/mine", func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Query().Get("data")
		p2p.Mine([]byte(data))
		w.Write([]byte("Block mined!"))
	})
	http.HandleFunc("/", p2p.Root)
	http.ListenAndServe(":8080", nil)
}

func (p2p *P2P) Stop() {
	http.DefaultServeMux = http.NewServeMux()
}
