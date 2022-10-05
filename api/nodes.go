package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	blchain "github.com/meliora019/blockchain/blockchain"
)

func Nodes(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(405)
		res.Write([]byte("405 method not allowed\n"))
		return
	}

	var body map[string][]string

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(400)
		return
	}

	blockchain := blchain.Get()

	for _, node := range body["nodes"] {
		err := blockchain.RegisterNode(node)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	resp := map[string]interface{}{
		"message":    "New nodes have been added",
		"totalNodes": blockchain.Nodes(),
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(500)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(201)
	res.Write(respJson)
}
