package pool

import "time"

// The object pool creational design pattern is used to prepare and keep
// multiple instances according to the demand expectation.

type TXpool chan *TXhandler

func New(total int) TXpool {
	p := make(TXpool, total)

	for i := 0; i < total; i++ {
		p <- new(TXhandler)
	}

	return p
}

type TXhandler struct {
	TX []byte
}

func (h *TXhandler) AddTX(tx []byte) {
	h.TX = tx
	time.Sleep(1 * time.Millisecond)
	block.AddTX(h.TX)
	h.TX = make([]byte, 0)
}
