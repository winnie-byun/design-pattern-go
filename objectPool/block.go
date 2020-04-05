package pool

var block Block

type Block struct {
	TX []byte
}

func (b *Block) AddTX(tx []byte) {
	b.TX = append(b.TX, tx...)
}

func (b *Block) GetBlock() string {
	return string(b.TX)
}
