package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPageInfo_HasMore(t *testing.T) {

	t.Run("has more pages", func(t *testing.T) {
		pi := &PageInfo{
			Current: 2,
			Next:    4,
		}
		assert.True(t, pi.HasMore())
	})

	t.Run("no more pages", func(t *testing.T) {
		pi := &PageInfo{
			Current: 3,
			Next:    3,
		}
		assert.False(t, pi.HasMore())
	})

}
