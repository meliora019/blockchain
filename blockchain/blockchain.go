package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"time"
)

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
}

type Block struct {
	Index        uint32
	Timestamp    time.Time
	Transactions []Transaction
	Proof        uint64
	PreviousHash string
}

type blockchain struct {
	chain               []Block
	currentTransactions []Transaction
	nodes               map[string]struct{}
}

var blockchainPointer *blockchain

// constructor
func New() *blockchain {
	b := new(blockchain)

	b.nodes = make(map[string]struct{})
	b.NewBlock("1", 100)

	return b
}

func init() {
	blockchainPointer = New()
}

func Get() *blockchain {
	return blockchainPointer
}

// Getter for chain
func (b *blockchain) FullChain() []Block {
	return b.chain
}

// Getter for currentTransactions
func (b *blockchain) CurrentTransactions() []Transaction {
	return b.currentTransactions
}

// Getter for nodes
func (b *blockchain) Nodes() []string {
	var nodes []string

	for node := range b.nodes {
		nodes = append(nodes, node)
	}

	return nodes
}

func (b *blockchain) NewBlock(previousHash string, proof uint64) Block {
	var block Block

	block.Index = uint32(len(b.chain)) + 1
	block.Timestamp = time.Now().UTC()
	block.Transactions = b.currentTransactions
	block.Proof = proof
	if previousHash == "" {
		i := len(b.chain) - 1
		block.PreviousHash = Hash(b.chain[i])
	} else {
		block.PreviousHash = previousHash
	}

	// Empty list of current transactions
	b.currentTransactions = nil

	b.chain = append(b.chain, block)
	return block
}

func (b *blockchain) NewTransaction(transaction Transaction) uint32 {
	b.currentTransactions = append(b.currentTransactions, transaction)

	lastBlock := b.LastBlock()

	return lastBlock.Index + 1
}

func (b *blockchain) LastBlock() Block {
	i := len(b.chain) - 1
	return b.chain[i]
}

func (b *blockchain) RegisterNode(address string) error {
	parsed, err := url.ParseRequestURI(address)
	if err != nil {
		return err
	}

	b.nodes[parsed.Host] = struct{}{}

	return nil
}

// Kinda "static" method
func Hash(block Block) string {
	blockJson, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}

	sha256 := sha256.New()
	sha256.Write([]byte(string(blockJson)))
	sum := sha256.Sum(nil)

	hash := hex.EncodeToString(sum[:])

	return hash
}
