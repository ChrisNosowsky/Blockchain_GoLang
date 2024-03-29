package main

import (
  "fmt"
  "strconv"
  "github.com/tensor-programming/golang-blockchain/blockchain"
)

func main() {
  chain := blockchain.InitBlockChain()

  chain.AddBlock("First Block After Genesis")
  chain.AddBlock("Second Block After Genesis")
  chain.AddBlock("Third Block After Genesis")

  for _, block := range chain.Blocks {
    fmt.Printf("Previous Hash: %x\n", block.PrevHash)
    fmt.Printf("Data in Block: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)
    fmt.Printf("\n")

    pow := blockchain.NewProof(block)
    fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
    fmt.Println()

  }
}
