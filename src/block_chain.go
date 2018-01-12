package main

// Blockchain represent the main structure of our BlockChain
type Blockchain struct {
	blocks []*Block
}

// AddBlock is a function to add a block to our BlockChain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewGenesisBlock is a function to initiate our BlockChain
func NewGenesisBlock() *Block {
	return NewBlock("Blockchain initiated !", []byte{})
}

// NewBlockchain is the init function of our BlockChain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
