package main

import (
	"fmt"

	c "github.com/filipovi/blockchain/chain"
)

func main() {
	c := c.New()
	c.Add("new 2")
	c.Add("new 3")

	fmt.Println("hello blockchain!")
	fmt.Println(c.Blocks)
}
