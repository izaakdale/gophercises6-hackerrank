package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	s := "middle-Outz"

	fmt.Print(cipher(s, 2) + "\n")

	s1 := "oneTwoThree"
	s2 := "saveChangesInTheEditor"
	s3 := ""

	countWordsInString(s1)
	countWordsInString(s2)
	countWordsInString(s3)
}

func cipher(s string, k int) string {
	cipherMap := rotateAlphabet(k)
	ret := ""
	for _, v := range s {
		upper := false
		if unicode.IsUpper(v) {
			upper = true
			v = unicode.ToLower(v)
		}

		conv, found := cipherMap[v]
		if found {
			if upper {
				ret += strings.ToUpper(string(conv))
			} else {
				ret += string(conv)
			}
		} else {
			ret += string(v)
		}
		// reset flag for next char
		upper = false
	}
	return ret
}

func rotateAlphabet(k int) map[rune]rune {

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	cipherSlice := []rune(alphabet)

	// move first element to the end k times
	for i := 0; i < k; i++ {
		// get first element
		fe := cipherSlice[0]
		cipherSlice = append(cipherSlice[1:], fe)
	}
	// create new slice of runes which will be the keys in the map
	alpharune := []rune(alphabet)
	// cipher map will map alphabet rune, to ciphered rune
	cipherMap := make(map[rune]rune)
	for k, v := range cipherSlice {
		cipherMap[alpharune[k]] = v
	}
	return cipherMap
}

func countWordsInString(s string) {

	n := 0

	if len(s) > 0 {
		n++
	}

	for _, v := range s {
		if unicode.IsUpper(int32(v)) {
			n++
		}
	}
	fmt.Printf("The string %s has %d words\n", s, n)
}
