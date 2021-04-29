package maths

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestAbs(t *testing.T) {
    assert.Equal(t, 5, Abs(5))
    assert.Equal(t, 5, Abs(-5))
}

func TestMin(t *testing.T) {
    assert.Equal(t, 5, Min(5, 8))
}

func TestMax(t *testing.T) {
    assert.Equal(t, 8, Max(5, 8))
}
