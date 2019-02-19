package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateBoard(t *testing.T) {
	assert := assert.New(t)
	var counts [13]int
	for _, slot := range TheBoard {
		for _, pipe := range slot {
			switch pipe.(type) {
			case Terminal:
				terminal := pipe.(Terminal)
				assert.True(terminal >= 1 && terminal <= 12)
				counts[terminal]++
			case Connector:
				connector := pipe.(Connector)
				assert.Equal(connector.Target().Target(), connector)
			}
		}
	}

	for i := 1; i < 13; i++ {
		assert.Equal(counts[i], 2, "there are %d terminals with value %d", counts[i], i)
	}
}
