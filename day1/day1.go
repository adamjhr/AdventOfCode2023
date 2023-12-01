package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func P1() int {
	var first rune
	var last rune
	var result int
	foundFirstDigit := false
	iteration := 0

	r := bufio.NewReader(os.Stdin)
	for {
		c, _, err := r.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if unicode.IsDigit(c) {
			if foundFirstDigit {
				last = c
			} else {
				first = c
				last = c
				foundFirstDigit = true
			}
		} else if c == '\n' {
			if foundFirstDigit {
				number, err := strconv.Atoi(string(first) + string(last))
				if err != nil {
					log.Fatal(err)
				}
				result += number
			}
			foundFirstDigit = false
			iteration++
		}
	}

	return result
}

func toNum(val string) string {
	switch val {
	case "0", "zero", "orez":
		return "0"
	case "1", "one", "eno":
		return "1"
	case "2", "two", "owt":
		return "2"
	case "3", "three", "eerht":
		return "3"
	case "4", "four", "ruof":
		return "4"
	case "5", "five", "evif":
		return "5"
	case "6", "six", "xis":
		return "6"
	case "7", "seven", "neves":
		return "7"
	case "8", "eight", "thgie":
		return "8"
	case "9", "nine", "enin":
		return "9"
	}
	return ""
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func P2() int {
	result := 0
	r := bufio.NewReader(os.Stdin)

	regex, err := regexp.Compile("(zero)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|[0-9]")
	if err != nil {
		log.Fatal(err)
	}

	regexReverse, err := regexp.Compile("(orez)|(eno)|(owt)|(eerht)|(ruof)|(evif)|(xis)|(neves)|(thgie)|(enin)|[0-9]")
	if err != nil {
		log.Fatal(err)
	}

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		all := regex.FindAllString(line, -1)
		allReverse := regexReverse.FindAllString(reverseString(line), -1)
		if len(all) != 0 {
			number, err := strconv.Atoi(toNum(all[0]) + toNum(allReverse[0]))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(number)
			result += number
		}
	}

	return result
}
