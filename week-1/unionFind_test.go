package week1_test

import (
	"testing"

	week1 "github.com/lzh2nix/algorithms101/week-1"
	"github.com/stretchr/testify/assert"
)

func TestQuickFind(t *testing.T) {
	ufs := []week1.DynamicConnect{week1.NewQuickFind(10)}

	for _, uf := range ufs {
		uf.Union(4, 3)
		uf.Union(3, 8)
		uf.Union(6, 5)
		uf.Union(9, 4)
		uf.Union(2, 1)
		assert.True(t, uf.Connected(1, 2), "test conncted(1, 2)")
		assert.False(t, uf.Connected(0, 7), "test conncted(0, 7)")
		assert.True(t, uf.Connected(8, 9), "test connected(8,9)")
		uf.Union(5, 0)
		uf.Union(7, 2)
		uf.Union(6, 1)
		uf.Union(1, 0)
		assert.True(t, uf.Connected(0, 7), "test connected(0,7)")
		assert.False(t, uf.Connected(9, 7), "test connected(7,9)")

	}

}
