package main

import (
	"fmt"
	"io"
	"net/http"

	"encoding/json"

	"github.com/quantumshiro/perychain/pkg/blockchain"
)

func blocks(w http.ResponseWriter, r *http.Request) {
	bc := blockchain.NewBlockChain()

	json.NewEncoder(w).Encode(bc.Chain)
}

func mine(w http.ResponseWriter, r *http.Request) {
	bc := blockchain.NewBlockChain()
	// data: request body data
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	bc.AddBlock(data)

	json.NewEncoder(w).Encode(bc.Chain)

	// redirect
	http.Redirect(w, r, "/blocks", http.StatusSeeOther)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/blocks", blocks)
	http.HandleFunc("/mine", mine)
	http.ListenAndServe(":8080", nil)
}
