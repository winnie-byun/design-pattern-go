package singleton

import (
	"sync"

	pattern "github.com/winnie-byun/design-pattern-go"
)

// Singleton creational design pattern restricts
// the instantiation of a type to a single object.
// https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f

// using init() : not secure, limitation in arguments and return

var (
	once       sync.Once
	blockchain *pattern.Blockchain
)

func NewBlockchain() *pattern.Blockchain {
	once.Do(func() {
		blockchain = pattern.NewBlockChain()
	})

	return blockchain
}

func NewBlock(block []byte) {
	blockchain.NewBlock(block)
}

func GetBlockNumber() int {
	return blockchain.BlockNumber
}

func GetBlock(blockNumber int) []byte {
	return blockchain.GetBlock(blockNumber)
}
