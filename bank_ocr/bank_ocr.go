package bank_ocr

import (
	"log"
)

// From https://codingdojo.org/kata/BankOCR/

const NumbersPerRow = 9
const CharactersPerNumber = 3
const RowLength = NumbersPerRow*CharactersPerNumber + 1
const TotalLength = RowLength*3 + 1

var HashToNumber = map[int]string{
    57662: "0",
    42848: "1",
    34382: "2",
    52046: "3",
    44592: "4",
    49838: "5",
    55726: "6",
    42974: "7",
    58670: "8",
    52782: "9",
}

func OcrParse(raw string) string {
	if len(raw) != TotalLength {
		log.Panicf("Invalid input, found size %d, should be %d", len(raw), TotalLength)
	}

	var result = ""
	for i := 0; i < NumbersPerRow; i++ {
		var characters = getCharactersAtIndex(raw, i)
		var hash = getCharactersHash(characters)
		var number = HashToNumber[hash]
		if number == "" {
            result = result + "?"
        } else {
            result = result + number
        }
	}

	return result
}

func getCharactersAtIndex(raw string, index int) string {
	var indexRow0 = CharactersPerNumber * index
	var indexRow1 = RowLength + CharactersPerNumber*index
	var indexRow2 = RowLength*2 + CharactersPerNumber*index
	row0 := raw[indexRow0 : indexRow0+CharactersPerNumber]
	row1 := raw[indexRow1 : indexRow1+CharactersPerNumber]
	row2 := raw[indexRow2 : indexRow2+CharactersPerNumber]
	return row0 + row1 + row2
}

func getCharactersHash(raw string) int {
	var acc = 0
	for i, ch := range raw {
		acc += int(ch) << i
	}
	return acc
}
