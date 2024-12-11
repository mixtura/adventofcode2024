package main

import (
	"os"
	"strconv"
	"unicode"
)

func day3() {
	file, err := os.ReadFile("input03.txt")
	if err != nil {
		panic(err)
	}

	text := string(file)
	textRunes := []rune(text)
	mulPrefix := "mul("
	mulPrefixParsed := false
	mulEnabled := true
	doToken := "do()"
	dontToken := "don't()"
	commaParsed := false
	firstNumberRaw := ""
	secondNumberRaw := ""
	total := 0

	reset := func() {
		commaParsed = false
		mulPrefixParsed = false
		firstNumberRaw = ""
		secondNumberRaw = ""
	}

	for idx := 0; idx < len(textRunes); idx++ {
		c := textRunes[idx]

		if mulPrefixParsed {
			if unicode.IsDigit(c) {
				if commaParsed {
					secondNumberRaw += string(c)
				} else {
					firstNumberRaw += string(c)
				}

				continue
			}

			if c == ',' {
				commaParsed = true
				continue
			}

			if c == ')' {
				if firstNumberRaw == "" || secondNumberRaw == "" {
					reset()
					continue
				}

				firstNumber, err := strconv.Atoi(firstNumberRaw)
				if err != nil {
					panic(err)
				}

				secondNumber, err := strconv.Atoi(secondNumberRaw)
				if err != nil {
					panic(err)
				}

				if mulEnabled {
					total += firstNumber * secondNumber
				}
			}

			reset()
		}

		mulPrefixParsed = consumeToken(textRunes, mulPrefix, idx)

		if mulPrefixParsed {
			idx += len(mulPrefix) - 1
			continue
		}

		doParsed := consumeToken(textRunes, doToken, idx)

		if doParsed {
			mulEnabled = true
			idx += len(doToken) - 1
			continue
		}

		dontParsed := consumeToken(textRunes, dontToken, idx)

		if dontParsed {
			mulEnabled = false
			idx += len(dontToken) - 1
			continue
		}
	}

	println(total)
}

func consumeToken(text []rune, token string, startIdx int) bool {
	for doIdx, doCh := range token {
		if text[startIdx+doIdx] != doCh {
			return false
		}
	}

	return true
}
