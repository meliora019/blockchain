package blockchain

import (
	"testing"
	// "fmt"
)

func TestConsensus(t *testing.T) {
	blockchain := New()

	var transaction Transaction

	transaction = Transaction{"sender1", "recipient1", 0.34}
	blockchain.NewTransaction(transaction)
	transaction = Transaction{"sender2", "recipient2", 2.34}
	blockchain.NewTransaction(transaction)
	lastBlock := blockchain.LastBlock()
	proof := blockchain.ProofOfWork(lastBlock)
	blockchain.NewBlock("", proof)

	transaction = Transaction{"sender3", "recipien31", 0.09}
	blockchain.NewTransaction(transaction)
	transaction = Transaction{"sender2", "recipient2", 2.34}
	blockchain.NewTransaction(transaction)
	lastBlock = blockchain.LastBlock()
	proof = blockchain.ProofOfWork(lastBlock)
	blockchain.NewBlock("", proof)

	isValid := blockchain.validChain(blockchain.chain)

	if isValid != true {
		t.Error("Expected true, got", isValid)
	}

	// blockchain.RegisterNode("http://127.0.0.1:8080")
	// blockchain.RegisterNode("http://127.0.0.1:5000")
	// blockchain.RegisterNode("http://127.0.0.1:5001")
	//
	// conflict := blockchain.ResolveConflicts()
	// fmt.Println(conflict)

	// fmt.Printf("%+v\n", blockchain)
	// for i, chain := range blockchain.chain {
	//   fmt.Printf("%d: %+v\n", i, chain)
	// }
}
