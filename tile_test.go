package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTileReflection(t *testing.T) {
	assert := assert.New(t)

	assert.Len(Tiles, 9)

	for _, tile := range Tiles {
		for posIndex, to := range tile {
			assert.Equal(tile[to], posIndex)
		}
	}
}
