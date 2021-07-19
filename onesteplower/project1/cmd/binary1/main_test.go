package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"project1/lib/lib1"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, int64(4), lib1.DoesThing(2), "they should be equal")
}
