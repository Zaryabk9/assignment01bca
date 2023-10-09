package assignment01bca

import (
    "crypto/sha256"
    "fmt"
    "time"
	"math/rand"
)

// Define the structure for a block
type Block struct {
    Transaction  string
    Nonce        int
    PreviousHash string
    Timestamp    int64
    CurrentHash         string
}

// CreateHash calculates the hash of a block
func CalculateHash(block *Block) string {
	data := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// NewBlock creates a new block
func NewBlock(transaction string, previousHash string) *Block {
	rand.Seed(time.Now().UnixNano()) 
    nonce := rand.Intn(1000000)      
    block := &Block{
        Transaction:  transaction,
        Nonce:        nonce,
        PreviousHash: previousHash,
        Timestamp:    time.Now().Unix(), 
    }

    // Calculate the block's hash using CalculateHash
    block.CurrentHash= CalculateHash(block)
	return block
    
}

// DisplayBlocks prints all the blocks in a nice format
func DisplayBlocks(blocks []*Block) {
    fmt.Println("Displaying Blocks:")
    for _, block := range blocks {
        fmt.Println("Transaction:", block.Transaction)
        fmt.Println("Nonce:", block.Nonce)
        fmt.Println("Previous Hash:", block.PreviousHash)
        fmt.Println("Current Hash:", block.CurrentHash)
        fmt.Println()
    }
}

// ChangeBlock changes the transaction of a given block reference
func ChangeBlock(block *Block, newTransaction string) {
    // Update the current block's transaction
    block.Transaction = newTransaction
	
	block.CurrentHash=CalculateHash(block)
}

// VerifyChain checks the integrity of the blockchain
func VerifyChain(blocks []*Block) bool {
    for i := 1; i < len(blocks); i++ {
        currentBlock := blocks[i]
        previousBlock := blocks[i-1]

        
        if currentBlock.PreviousHash != previousBlock.CurrentHash {
            return false // Chain has been tampered with
        }
        
        //currentBlock.CalculateHash()
        if currentBlock.CurrentHash != CalculateHash(currentBlock) {
            return false 
        }
    }
    return true // Blockchain is valid
}
