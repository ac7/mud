package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Note that this test fails intermittently: approximately 1 out of every
// 18446744073709551616 runs
func TestNewID(t *testing.T) {
	id1 := newID()
	id2 := newID()
	assert.NotEqual(t, id1, id2)
}
