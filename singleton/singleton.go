package singleton

import (
	"sync"

	"github.com/winnie-byun/design-pattern-go/common"
)

// Singleton creational design pattern restricts
// the instantiation of a type to a single object.
// https://medium.com/golang-issue/how-singleton-pattern-works-with-golang-2fdd61cd5a7f

// using init() : not secure, limitation in arguments and return

var (
	once       sync.Once
	blockchain *common.Blockchain
)

func NewBlockchain() *common.Blockchain {
	once.Do(func() {
		blockchain = common.NewBlockChain()
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
