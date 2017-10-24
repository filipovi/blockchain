package chain

import (
	"fmt"

	"github.com/filipovi/blockchain/block"
)

// Chain is a struct who contains all the blocks
type Chain struct {
	Blocks []block.Block `json:"blocs"`
}

const genesis = "genesis block"

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
func (c *Chain) Add(b block.Block) error {
	l := len(c.Blocks)
	if b.Index != l {
		return fmt.Errorf("block not valid, index")
	}
	lb := c.GetLatestBlock()
	if b.PHash != lb.Hash {
		return fmt.Errorf("block not valid, pHash")
	}
	if b.Hash != b.CalculateHash() {
		return fmt.Errorf("block not valid, hash")
	}

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

	gb := block.New(genesis, "0", 0)
	blocks[0] = gb

	c.Blocks = blocks

	return c
}
