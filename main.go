package main

import (
	"fmt"

	"github.com/filipovi/blockchain/block"
	"github.com/filipovi/blockchain/chain"
)

func main() {
	c := chain.New()

	lb := c.GetLatestBlock()
	b := block.New("First block", lb.Hash, lb.Index+1)
	c.Add(b)

	lb = c.GetLatestBlock()
	b2 := block.New("Second block", lb.Hash, lb.Index+1)
	c.Add(b2)

	b3 := block.Block{
		Hash:  "0",
		Data:  "fake",
		PHash: "0",
		Index: 1,
	}
	err := c.Add(b3)
	fmt.Println(err)

	fmt.Println("hello blockchain!")
	for _, b := range c.Blocks {
		fmt.Println(b)
	}
}
