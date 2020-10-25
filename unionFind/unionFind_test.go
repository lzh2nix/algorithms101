package unionFind_test

import (
	"testing"

	"github.com/lzh2nix/algorithms101/unionFind"
	"github.com/stretchr/testify/assert"
)

func test(uf unionFind.DynamicConnect, t *testing.T, tc string) {

	uf.Union(4, 3)
	uf.Union(3, 8)
	uf.Union(6, 5)
	uf.Union(9, 4)
	uf.Union(2, 1)
	assert.True(t, uf.Connected(1, 2), "test conncted(1, 2) for "+tc)
	assert.False(t, uf.Connected(0, 7), "test conncted(0, 7) for "+tc)
	assert.True(t, uf.Connected(8, 9), "test connected(8,9) for "+tc)
	uf.Union(5, 0)
	uf.Union(7, 2)
	uf.Union(6, 1)
	uf.Union(1, 0)
	assert.True(t, uf.Connected(0, 7), "test connected(0,7) for "+tc)
	assert.False(t, uf.Connected(9, 7), "test connected(7,9) for "+tc)
}
func TestQuickFind(t *testing.T) {
	test(unionFind.NewQuickFind(10), t, "quickFind")

}

func TestQuickUnion(t *testing.T) {
	test(unionFind.NewQuickUnion(10), t, "quickUnion")
}

func TestQuickUnionLazy(t *testing.T) {
	test(unionFind.NewQuickUnionLazy(10), t, "quickUnionLazy")
}

func TestQuickUnionWeight(t *testing.T) {
	test(unionFind.NewQuickUnionWeight(10), t, "quickUnionWeight")
}
