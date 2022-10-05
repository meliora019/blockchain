package blockchain

import (
	"testing"
	"time"
)

func TestValidProof(t *testing.T) {
	blockchain := New()

	boolean := blockchain.ValidProof(200, 0, "caf8e151d7b85676ff16e44fa79115b4a5e2f5ab096764d02152c83d0c03ee1b")

	if boolean != false {
		t.Error("Expected false, got", boolean)
	}

	boolean = blockchain.ValidProof(200, 29, "caf8e151d7b85676ff16e44fa79115b4a5e2f5ab096764d02152c83d0c03ee1b")

	if boolean != false {
		t.Error("Expected false, got", boolean)
	}

	boolean = blockchain.ValidProof(200, 4781, "caf8e151d7b85676ff16e44fa79115b4a5e2f5ab096764d02152c83d0c03ee1b")

	if boolean != true {
		t.Error("Expected true, got", boolean)
	}
}

func TestProofOfWork(t *testing.T) {
	blockchain := New()

	var block Block

	block.Index = 2
	block.Transactions = append(block.Transactions, Transaction{"sender3", "recipient3", 0.09})
	block.Transactions = append(block.Transactions, Transaction{"sender2", "recipient2", 2.34})
	block.Proof = 200
	block.PreviousHash = "8312748868ad0bbf2e6128c1f34cf6aa3b13dd73ae0d98dced8297edbd005d63"
	block.Timestamp = time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC)
	// Hash(block) = "caf8e151d7b85676ff16e44fa79115b4a5e2f5ab096764d02152c83d0c03ee1b"

	proof := blockchain.ProofOfWork(block)

	if proof != 4781 {
		t.Error("Expected 4781, got", proof)
	}
}
