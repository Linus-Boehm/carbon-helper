package carbon_helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/uniplaces/carbon"
	"testing"
	"time"
)

func Test_Add(t *testing.T) {
	ti := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	c := carbon.NewCarbon(ti)
	ch := New(c)
	ch.Add("day", 1)
	ex := time.Date(2009, time.November, 11, 23, 0, 0, 0, time.UTC)
	assert.Equal(t, ex, ch.Carbon.Time)
}
