package bank_ocr

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func readSample(t *testing.T, filename string) string {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        t.Fatal(err)
    }
    return string(bytes)
}

func TestParse000000000(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_000000000.txt")
    assert.Equal(t, " _ | ||_|", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "000000000", OcrParse(sample))
}

func TestParse111111111(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_111111111.txt")
    assert.Equal(t, "     |  |", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "111111111", OcrParse(sample))
}

func TestParse222222222(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_222222222.txt")
    assert.Equal(t, " _  _||_ ", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "222222222", OcrParse(sample))
}

func TestParse333333333(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_333333333.txt")
    assert.Equal(t, " _  _| _|", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "333333333", OcrParse(sample))
}

func TestParse444444444(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_444444444.txt")
    assert.Equal(t, "   |_|  |", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "444444444", OcrParse(sample))
}

func TestParse555555555(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_555555555.txt")
    assert.Equal(t, " _ |_  _|", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "555555555", OcrParse(sample))
}

func TestParse666666666(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_666666666.txt")
    assert.Equal(t, " _ |_ |_|", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "666666666", OcrParse(sample))
}

func TestParse777777777(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_777777777.txt")
    assert.Equal(t, " _   |  |", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "777777777", OcrParse(sample))
}

func TestParse888888888(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_888888888.txt")
    assert.Equal(t, " _ |_||_|", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "888888888", OcrParse(sample))
}

func TestParse999999999(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_999999999.txt")
    assert.Equal(t, " _ |_| _|", getCharactersAtIndex(sample, 0))
    assert.Equal(t, "999999999", OcrParse(sample))
}

func TestParse123456789(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_123456789.txt")
    assert.Equal(t, "123456789", OcrParse(sample))
}

func TestParse000000051(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_000000051.txt")
    assert.Equal(t, "000000051", OcrParse(sample))
}

func TestParse49006771_(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_49006771_.txt")
    assert.Equal(t, "49006771?", OcrParse(sample))
}

func TestParse1234_678_(t *testing.T) {
    sample := readSample(t, "bank_ocr_test_1234_678_.txt")
    assert.Equal(t, "1234?678?", OcrParse(sample))
}

func TestGetCharacterHash(t *testing.T) {
}
