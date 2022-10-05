package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	transactions []string
	previoushash []byte
	currenthash  []byte
}

func main() {
	a := []string{"Tehami sent 100 coins to Ahmad"}
	b := blocks(a, []byte{})
	fmt.Println("-------First block-------")
	listblock(b)

	x := []string{"Ali sent 200 coins to Abdullah"}
	y := blocks(x, b.currenthash)
	fmt.Println("-------New block-------")
	listblock(y)
}

func blocks(transactions []string, previoushash []byte) *block {
	return &block{
		transactions: transactions,
		previoushash: previoushash,
		currenthash:  calculatehash(transactions, previoushash),
	}
}

func calculatehash(transactions []string, previoushash []byte) []byte {
	x := append(previoushash)
	for transaction := range transactions {
		x = append(x, string((transaction))...)
	}
	Hash := sha256.Sum256(x)
	return Hash[:]
}

func listblock(Block *block) {
	fmt.Printf("Transaction: %s\n", Block.transactions[:])
	fmt.Printf("previous bloch hash: %x\n", Block.previoushash)
	fmt.Printf("current block hash: %x\n\n", Block.currenthash)
}
