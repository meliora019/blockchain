package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	blchain "github.com/meliora019/blockchain/blockchain"
)

func Consensus(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		res.WriteHeader(405)
		res.Write([]byte("405 method not allowed"))
		return
	}

	blockchain := blchain.Get()

	replaced := blockchain.ResolveConflicts()

	type response struct {
		Message string `json:"message"`
		Chain   []blchain.Block
	}

	var resp response

	if replaced {
		resp.Message = "Our chain was replaced"
		resp.Chain = blockchain.FullChain()
	} else {
		resp.Message = "Our chain is authoritative"
		resp.Chain = blockchain.FullChain()
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(500)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(respJson)
}
