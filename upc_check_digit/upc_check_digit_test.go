package upc_check_digit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpcCheckDigit(t *testing.T) {
	assert.Equal(t, 4, UpcCheckDigit(4210000526))
	assert.Equal(t, 2, UpcCheckDigit(3600029145))
	assert.Equal(t, 4, UpcCheckDigit(12345678910))
	assert.Equal(t, 0, UpcCheckDigit(1234567))
}

func TestNthDigit(t *testing.T) {
	assert.Equal(t, 5, nthDigit(3600029145, 0))
	assert.Equal(t, 4, nthDigit(3600029145, 1))
	assert.Equal(t, 1, nthDigit(3600029145, 2))
	assert.Equal(t, 9, nthDigit(3600029145, 3))
	assert.Equal(t, 2, nthDigit(3600029145, 4))
	assert.Equal(t, 0, nthDigit(3600029145, 5))
	assert.Equal(t, 0, nthDigit(3600029145, 20))
}
