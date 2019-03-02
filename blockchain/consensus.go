package blockchain

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "fmt"
)

func (b *blockchain) validChain(chain []Block) bool {
  lastBlock := chain[0]

  for currentIndex := 1; currentIndex < len(chain); currentIndex++ {
    block := chain[currentIndex]

    if block.PreviousHash != Hash(lastBlock) {
      return false
    }

    if !b.ValidProof(lastBlock.Proof, block.Proof, Hash(lastBlock)) {
      return false
    }

    lastBlock = block
  }

  return true
}

func (b *blockchain) ResolveConflicts() bool {
  neighbours := b.nodes
  var newChain []Block
  maxLength := len(b.chain)

  type body struct {
    Chain []Block
    Length int
  }

  for node := range neighbours {
    url := fmt.Sprintf("http://%s/chain", node)
    response, err := http.Get(url)
    if err != nil {
      fmt.Println(err)
      continue
    }
    contents, err := ioutil.ReadAll(response.Body)
    if err != nil {
      fmt.Println(err)
      continue
    }

    var bd body

    err = json.Unmarshal(contents, &bd)
    if err != nil {
      fmt.Println(err)
      continue
    }

    if response.StatusCode == 200 {
      length := bd.Length
      chain := bd.Chain

      if length > maxLength && b.validChain(chain) {
        maxLength = length
        newChain = chain
      }
    }

    response.Body.Close()
  }

  if len(newChain) != 0 {
    b.chain = newChain
    return true
  }

  return false
}
