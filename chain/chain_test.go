package chain

import (
	"testing"

	"github.com/filipovi/blockchain/block"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()

	assert.Equal(t, 1, len(c.Blocks))
	assert.Equal(t, genesis, c.Blocks[0].Data)
	assert.Equal(t, 0, c.Blocks[0].Index)
	assert.Equal(t, "0", c.Blocks[0].PHash)
}

func TestGetLastBlock(t *testing.T) {
	c := New()

	lb := c.GetLatestBlock()
	assert.Equal(t, genesis, lb.Data)
	assert.Equal(t, 0, lb.Index)
	assert.Equal(t, "0", lb.PHash)
}

func TestAdd(t *testing.T) {
	c := New()
	lb := c.GetLatestBlock()
	ph := lb.Hash
	b := block.New("new block", lb.Hash, lb.Index+1)
	err := c.Add(b)
	assert.Nil(t, err)

	lb = c.GetLatestBlock()
	assert.Equal(t, "new block", lb.Data)
	assert.Equal(t, 1, lb.Index)
	assert.Equal(t, ph, lb.PHash)
}

func testAddNotValid(t *testing.T) {
	c := New()
	lb := c.GetLatestBlock()

	b := block.New("new block", lb.Hash, 127)
	err := c.Add(b)
	assert.Equal(t, "block not valid, index", err)

	b = block.New("new block", "xxxx", lb.Index+1)
	err = c.Add(b)
	assert.Equal(t, "block not valid, pHash", err)

	b = block.New("new block", lb.Hash, lb.Index+1)
	b.Hash = "xxxxx"
	err = c.Add(b)
	assert.Equal(t, "block not valid, hash", err)
}

func testGet(t *testing.T) {
	c := New()
	lb := c.GetLatestBlock()
	ph := lb.Hash
	b := block.New("new block", lb.Hash, lb.Index+1)
	err := c.Add(b)
	assert.Nil(t, err)

	nb, err := c.Get(b.Hash)
	assert.Nil(t, err)
	assert.Equal(t, "new block", nb.Data)
	assert.Equal(t, 1, nb.Index)
	assert.Equal(t, ph, nb.PHash)
}

func testGetError(t *testing.T) {
	c := New()
	lb := c.GetLatestBlock()
	b := block.New("new block", lb.Hash, lb.Index+1)
	err := c.Add(b)
	assert.Nil(t, err)

	_, err = c.Get("xxxx")
	assert.Equal(t, "block xxxx not found in the chain", err)
}
