package blockchain

import (
  "testing"
  "time"
)

func TestNew(t *testing.T) {
  // blockchain := New()

  if len(blockchainPointer.currentTransactions) != 0 {
    t.Error("Expected 0, got", len(blockchainPointer.currentTransactions))
  }
  if len(blockchainPointer.chain) != 1 {
    t.Error("Expected 1, got", len(blockchainPointer.chain))
  }
  if blockchainPointer.chain[0].Index != 1 {
    t.Error("Expected 1, got", blockchainPointer.chain[0].Index)
  }
  if len(blockchainPointer.chain[0].Transactions) != 0 {
    t.Error("Expected 0, got", len(blockchainPointer.chain[0].Transactions))
  }
  if blockchainPointer.chain[0].PreviousHash != "1" {
    t.Error("Expected 1, got", blockchainPointer.chain[0].PreviousHash)
  }
  if blockchainPointer.chain[0].Proof != 100 {
    t.Error("Expected 100, got", blockchainPointer.chain[0].Proof)
  }
  if len(blockchainPointer.nodes) != 0 {
    t.Error("Expected 0, got", len(blockchainPointer.nodes))
  }
}

func TestNewTransaction(t *testing.T) {
  blockchain := New()

  var transaction Transaction

  transaction = Transaction{"sender1", "recipient1", 0.34}
  blockchain.NewTransaction(transaction)
  transaction = Transaction{"sender2", "recipient2", 2.34}
	blockchain.NewTransaction(transaction)

  if len(blockchain.currentTransactions) != 2 {
    t.Error("Expected 2, got", len(blockchain.currentTransactions))
  }
  if blockchain.currentTransactions[0].Sender != "sender1" {
    t.Error("Expected sender1, got", blockchain.currentTransactions[0].Sender)
  }
  if blockchain.currentTransactions[1].Recipient != "recipient2" {
    t.Error("Expected recipient2, got", blockchain.currentTransactions[1].Recipient)
  }
  if blockchain.currentTransactions[1].Amount != 2.34 {
    t.Error("Expected 2.34, got", blockchain.currentTransactions[1].Amount)
  }
}

func TestNewBlockAndLastBlock(t *testing.T) {
  blockchain := New()

  var transaction Transaction

  transaction = Transaction{"sender1", "recipient1", 0.34}
  blockchain.NewTransaction(transaction)
  transaction = Transaction{"sender2", "recipient2", 2.34}
  blockchain.NewTransaction(transaction)
  blockchain.NewBlock("", 200);

  transaction = Transaction{"sender3", "recipien31", 0.09}
  blockchain.NewTransaction(transaction)
  transaction = Transaction{"sender2", "recipient2", 2.34}
  blockchain.NewTransaction(transaction)
  blockchain.NewBlock("", 400);

  if len(blockchain.currentTransactions) != 0 {
    t.Error("Expected 0, got", len(blockchain.currentTransactions))
  }
  if len(blockchain.chain) != 3 {
    t.Error("Expected 3, got", len(blockchain.chain))
  }
  if len(blockchain.chain[1].Transactions) != 2 {
    t.Error("Expected 2, got", len(blockchain.chain[1].Transactions))
  }
  if blockchain.chain[1].Proof != 200 {
    t.Error("Expected 200, got", blockchain.chain[1].Proof)
  }
  if len(blockchain.chain[2].Transactions) != 2 {
    t.Error("Expected 2, got", len(blockchain.chain[2].Transactions))
  }
  if blockchain.chain[2].Proof != 400 {
    t.Error("Expected 400, got", blockchain.chain[2].Proof)
  }

  lastBlock := blockchain.LastBlock()
  if lastBlock.Index != 3 {
    t.Error("Expected 3, got", lastBlock.Index)
  }
  if len(lastBlock.Transactions) != 2 {
    t.Error("Expected 2, got", len(lastBlock.Transactions))
  }
  if lastBlock.Proof != 400 {
    t.Error("Expected 400, got", lastBlock.Proof)
  }
}

func TestRegiserNode(t *testing.T) {
  blockchainPointer.RegisterNode("http://192.168.0.5:5000")
  blockchainPointer.RegisterNode("http://192.168.0.5:5000")
  blockchainPointer.RegisterNode("http://192.168.0.5:5001")
  err := blockchainPointer.RegisterNode("http://192.168.0.4:5000")

  if err != nil {
    t.Error("Expected nil, got", err)
  }

  err = blockchainPointer.RegisterNode("sdfsdfsdf")
  if err == nil {
    t.Error("Expected some error, got", err)
  }

  if len(blockchainPointer.nodes) != 3 {
    t.Error("Expected 3, got", len(blockchainPointer.nodes))
  }
  if _, ok := blockchainPointer.nodes["192.168.0.5:5000"]; !ok {
    t.Error("Expected true, got", ok)
  }
  if _, ok := blockchainPointer.nodes["192.168.0.5:5001"]; !ok {
    t.Error("Expected true, got", ok)
  }
  if _, ok := blockchainPointer.nodes["192.168.0.4:5000"]; !ok {
    t.Error("Expected true, got", ok)
  }
  if _, ok := blockchainPointer.nodes["192.168.0.5:5002"]; ok {
    t.Error("Expected false, got", ok)
  }
}

func TestHash(t *testing.T) {
  var block Block

  block.Index = 3
  block.Transactions = append(block.Transactions, Transaction{"sender3", "recipient3", 0.09})
  block.Transactions = append(block.Transactions, Transaction{"sender2", "recipient2", 2.34})
  block.Proof = 400
  block.PreviousHash = "8312748868ad0bbf2e6128c1f34cf6aa3b13dd73ae0d98dced8297edbd005d63"
  block.Timestamp = time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC)

  hash := Hash(block)

  if hash != "729f127ed57c6d9cbdbbac8da84207ae112b8d81196cbd86c14f002cfd760e5f" {
    t.Error("Expected 729f127ed57c6d9cbdbbac8da84207ae112b8d81196cbd86c14f002cfd760e5f, got", hash)
  }
}
