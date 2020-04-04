package pattern

type Blockchain struct {
	BlockNumber int
	Blocks      [][]byte
}

func NewBlockChain() *Blockchain {
	return &Blockchain{}
}

// TODO: need lock
func (bc *Blockchain) NewBlock(block []byte) {
	bc.BlockNumber++
	bc.Blocks = append(bc.Blocks, block)
}

func (bc *Blockchain) GetBlockNumber() int {
	return bc.BlockNumber
}

func (bc *Blockchain) GetBlock(blockNumber int) []byte {
	return bc.Blocks[blockNumber]
}
