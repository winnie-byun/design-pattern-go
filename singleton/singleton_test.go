package singleton

import (
	"log"
	"testing"
)

func TestSingleton(t *testing.T) {
	b1 := NewBlockchain()
	log.Println(b1.GetBlockNumber())
	b1.NewBlock([]byte("This is the first block"))
	log.Println(b1.GetBlockNumber())

	b2 := NewBlockchain()
	log.Println(b2.GetBlockNumber())
	b2.NewBlock([]byte("This is second block"))
	log.Println(b2.GetBlockNumber())

	log.Println(string(b2.GetBlock(0)))
	log.Println(string(b1.GetBlock(1)))
}

func TestStatic(t *testing.T) {
	NewBlockchain()

	log.Println(GetBlockNumber())
	NewBlock([]byte("This is the first block"))
	log.Println(GetBlockNumber())
	NewBlock([]byte("This is second block"))
	log.Println(GetBlockNumber())

	log.Println(string(GetBlock(0)))
	log.Println(string(GetBlock(1)))
}
