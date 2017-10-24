package main

import (
	"fmt"

	"github.com/filipovi/blockchain/chain"
)

func main() {
	c := chain.New()

	fmt.Println("hello blockchain!")
	for _, b := range c.Blocks {
		fmt.Println(b)
	}
}
