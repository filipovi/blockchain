package chain

import (
	"github.com/filipovi/blockchain/block"
)

type Chain struct {
	Blocks []block.Block `json:"blocs"`
}

func (c Chain) getLatestBlock() block.Block {
	return c.Blocks[len(c.Blocks)-1]
}

func (c *Chain) Add(data string) error {
	l := len(c.Blocks)
	b := block.New(data)
	lb := c.getLatestBlock()
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
