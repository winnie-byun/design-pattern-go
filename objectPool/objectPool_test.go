package pool

import (
	"fmt"
	"log"
	"testing"
)

func TestPool(t *testing.T) {
	p := (chan *TXhandler)(New(5))

	for i := 0; i <= 20; i++ {
		select {
		case txHandler := <-p:
			tx := fmt.Sprintf("%d sent %d $%d", i, 100-i, i)
			txHandler.addTX(tx)
		default:
			log.Println("TxPool is full")
		}
	}

	log.Println(string(block.TX))
}
