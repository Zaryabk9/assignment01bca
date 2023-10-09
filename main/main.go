package main

import (
	"github.com/Zaryabk9/assignment01bca/assignment01bca"
    "fmt"
)

func main() {
    
	//function calling
    blocks := []*assignment01bca.Block{}

   
    transaction1 := "Alice to Bob"
    previousHash1 := "genesis_hash"

    blocks = append(blocks, assignment01bca.NewBlock(transaction1, previousHash1))
   
    assignment01bca.DisplayBlocks(blocks)

    newTransaction := "Charlie to Dave"
    assignment01bca.ChangeBlock(blocks[0], newTransaction)
  
    assignment01bca.DisplayBlocks(blocks)

    isValid := assignment01bca.VerifyChain(blocks)
    if isValid {
        fmt.Println("Blockchain is valid.")
    } else {
        fmt.Println("Blockchain has been tampered with.")
    }
}
