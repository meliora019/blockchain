package api

import (
  "net/http"
  "encoding/json"
  blchain "blockchain/blockchain"
  "blockchain/node"
  "fmt"
)

func Mine(res http.ResponseWriter, req *http.Request) {
  if req.Method != "GET" {
    res.WriteHeader(405)
    res.Write([]byte("405 method not allowed"))
    return
  }

  type response struct {
    Success string `json:"success"`
    Message string `json:"message"`
    Index uint32
    Transactions []blchain.Transaction
    Proof uint64
    PreviousHash string
  }

  blockchain := blchain.Get()
  node := node.Get()

  lastBlock := blockchain.LastBlock()

  proof := blockchain.ProofOfWork(lastBlock)

  var transaction blchain.Transaction

  // We have to get a reward for found confirmation.
  // Sender "0" means that the node earned coin.
  transaction.Sender = "0"
  transaction.Recipient = node
  transaction.Amount = 1

  blockchain.NewTransaction(transaction)

  block := blockchain.NewBlock("", proof)

  var resp response
  resp.Success = "1"
  resp.Message = "New block created"
  resp.Index = block.Index
  resp.Transactions = block.Transactions
  resp.Proof = block.Proof
  resp.PreviousHash = block.PreviousHash

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
