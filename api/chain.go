package api

import (
	"encoding/json"
	"net/http"

	blchain "github.com/meliora019/blockchain/blockchain"
)

func Chain(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		res.WriteHeader(405)
		res.Write([]byte("405 method not allowed"))
		return
	}

	type response struct {
		Chain  []blchain.Block
		Length int
	}

	blockchain := blchain.Get()

	var resp response
	resp.Chain = blockchain.FullChain()
	resp.Length = len(blockchain.FullChain())

	respJson, err := json.Marshal(resp)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(respJson)
}
