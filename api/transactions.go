package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	blchain "github.com/meliora019/blockchain/blockchain"
)

func Transactions(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(405)
		res.Write([]byte("405 method not allowed\n"))
		return
	}

	var transaction blchain.Transaction

	jsonDecoder := json.NewDecoder(req.Body)
	err := jsonDecoder.Decode(&transaction)
	if err != nil {
		fmt.Println(err)

		resp := make(map[string]string)
		resp["success"] = "0"
		resp["message"] = "Invalid data"

		respJson, err := json.Marshal(resp)
		if err != nil {
			fmt.Println(err)
			res.WriteHeader(500)
			return
		}

		res.WriteHeader(400)
		res.Write(respJson)
		return
	}

	// TODO: validate transaction.sender and transaction.recipient

	blockchain := blchain.Get()

	index := blockchain.NewTransaction(transaction)

	resp := make(map[string]string)
	resp["success"] = "1"
	resp["message"] = fmt.Sprintf("Transaction will be added to block %d", index)

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
