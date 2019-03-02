package blockchain

import (
  "fmt"
  "crypto/sha256"
  "encoding/hex"
)

func (b *blockchain) ProofOfWork(lastBlock Block) uint64 {
  var proof uint64
  proof = 0

  lastProof := lastBlock.Proof
  lastHash := Hash(lastBlock)

  for {
    if b.ValidProof(lastProof, proof, lastHash) {
      break
    }
    proof += 1
  }

  return proof
}

func (b *blockchain) ValidProof(lastProof, proof uint64, lastHash string) bool {
  guess := fmt.Sprintf("%d%d%s", lastProof, proof, lastHash)

  sha256 := sha256.New()
  sha256.Write([]byte(guess))
  sum := sha256.Sum(nil)

  hash := hex.EncodeToString(sum[:])

  first4 := hash[:4]

  return first4 == "0000"
}
