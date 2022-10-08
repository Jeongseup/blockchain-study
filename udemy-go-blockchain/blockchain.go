package main

import (
	"fmt"
	"log"
	"time"
)

// ============ Block ===============

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previouseHash string) *Block {
	b := new(Block)

	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previouseHash

	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp			%d\n", b.timestamp)
	fmt.Printf("nonce		    	%d\n", b.nonce)
	fmt.Printf("previous_hash		%s\n", b.previousHash)
	fmt.Printf("transactions		%s\n", b.transactions)
}

// ============ BlockChain ===============

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")

	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previouseHash string) *Block {
	b := NewBlock(nonce, previouseHash)
	bc.chain = append(bc.chain, b)

	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("Chain %d \n", i)
		block.Print()
	}
}

// ============ MAIN ===============

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()

	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()

	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()

}
