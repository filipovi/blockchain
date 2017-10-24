package block

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	b := New("data", "hash", 0)

	assert.Equal(t, "data", b.Data)
	assert.Equal(t, "hash", b.PHash)
	assert.Equal(t, 0, b.Index)
}

func TestCalculateHash(t *testing.T) {
	b := New("data", "hash", 0)
	hash := b.Hash

	b.CalculateHash()
	assert.Equal(t, hash, b.Hash)
}

func TestTimestamp(t *testing.T) {
	b := New("data", "hash", 0)

	const longForm = "2017-10-24 23:14:10.725419028 +0200 CEST m=+0.000322049"
	tb, _ := time.Parse(longForm, b.Timestamp)
	tm := time.Now()
	assert.True(t, tm.Sub(tb) > 0)

	b2 := New("new data", "new hash", 0)
	assert.NotEqual(t, b.Timestamp, b2.Timestamp)
}
