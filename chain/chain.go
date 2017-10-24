package chain

import (
	"fmt"

	"github.com/filipovi/blockchain/block"
)

// Chain is a struct who contains all the blocks
type Chain struct {
	Blocks []block.Block `json:"blocs"`
}

// GetLatestBlock return the last block of the chain
func (c Chain) GetLatestBlock() block.Block {
	return c.Blocks[len(c.Blocks)-1]
}

// Get return a block with the given hash
func (c Chain) Get(hash string) (block.Block, error) {
	for _, b := range c.Blocks {
		if b.Hash == hash {
			return b, nil
		}
	}

	var b block.Block

	return b, fmt.Errorf("block %s not found in the chain", hash)
}

// Add a new block at the end of the chain
func (c *Chain) Add(data string) error {
	l := len(c.Blocks)
	lb := c.GetLatestBlock()
	b := block.New(data)
	b.PHash = lb.Hash
	b.Index = l

	blocks := append(c.Blocks, b)
	c.Blocks = blocks

	return nil
}

// New create a new chain
func New() Chain {
	blocks := make([]block.Block, 1)

	c := Chain{
		Blocks: blocks,
	}

	b := block.New("init blockchain")
	blocks[0] = b

	c.Blocks = blocks

	return c
}
