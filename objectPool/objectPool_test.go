package pool

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p := New(5)

	for i := 0; i <= 100; i++ {
		select {
		case txHandler := <-p:
			tx := fmt.Sprintf("%d sent %d $%d\n", i, 100-i, i)
			txHandler.AddTX([]byte(tx))
		default:
			log.Println("TxPool is full")
			time.Sleep(5 * time.Millisecond)
		}
	}

	log.Println(string(block.TX))
}
