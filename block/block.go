package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block is a struct to manipulate block in blockchain
type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
	PHash     string `json:"p_hash"`
}

// CalculateHash return a sha256 sum for the Block
func (b Block) CalculateHash() string {
	input := strconv.Itoa(b.Index) + b.Timestamp + b.PHash + b.Data
	h := sha256.New()
	h.Write([]byte(input))

	return hex.EncodeToString(h.Sum(nil))
}

// New create a new block
func New(data, pHash string, index int) Block {
	b := Block{
		Timestamp: time.Now().String(),
		Data:      data,
		PHash:     pHash,
		Index:     index,
	}

	b.Hash = b.CalculateHash()
	return b
}
