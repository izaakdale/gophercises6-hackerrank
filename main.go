package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	s := "middle-Outz"
	fmt.Print(cipher(s, 2) + "\n")

	countWordsInString("oneTwoThree")
	countWordsInString("saveChangesInTheEditor")
	countWordsInString("")

	arr := []int{-4, 3, -9, 0, 4, 1}
	countArrayNumbers(arr)

	arr2 := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr2)

	timeConversion("05:01:00PM")
	timeConversion("05:01:00AM")

	findLonelyInt([]int{1, 2, 3, 4, 3, 2, 1})

	diagonalDifference([][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	})

}

func diagonalDifference(arr [][]int) {
	lrSum, rlSum := 0, 0

	for i := 0; i < len(arr); i++ {
		lrSum += arr[i][i]
		rlSum += arr[i][len(arr)-i-1]
	}

	var diff int
	if lrSum > rlSum {
		diff = lrSum - rlSum
	} else {
		diff = rlSum - lrSum
	}

	fmt.Println(diff)
}

func findLonelyInt(arr []int) {
	var lonelyInt int

	intFinder := make(map[int]bool)
	for _, v := range arr {
		_, found := intFinder[v]
		if !found {
			lonelyInt = v
			intFinder[v] = false
		} else {
			intFinder[v] = true
		}
	}
	fmt.Println(lonelyInt)
}

func timeConversion(s string) {
	var hh, mm, ss int
	var err error

	hh, err = strconv.Atoi(s[0:2])
	mm, err = strconv.Atoi(s[3:5])
	ss, err = strconv.Atoi(s[6:8])
	if err != nil {
		log.Fatal(err.Error())
	}

	if hh == 12 {
		if s[8:] == "AM" {
			hh = 0
		}
	} else {
		if s[8:] == "PM" {
			hh += 12
		}
	}

	fmt.Printf("%.2d:%.2d:%.2d\n", hh, mm, ss)
}

func miniMaxSum(arr []int32) {
	var sum, min, max int32

	for _, v := range arr {
		sum += v
	}
	// in order to compare min needs to be a high value
	max = 0
	min = sum
	for _, v := range arr {
		if sum-v < min {
			min = sum - v
		}
		if sum-v > max {
			max = sum - v
		}
	}

	fmt.Printf("%d %d\n", min, max)
}

func countArrayNumbers(arr []int) {
	var nPos, nNeg, nZero float32
	for _, v := range arr {
		if v > 0 {
			nPos++
		} else if v < 0 {
			nNeg++
		} else {
			nZero++
		}
	}
	size := float32(6)

	posProp := nPos / size
	negProp := nNeg / size
	zeroProp := nZero / size

	fmt.Printf("%.6f\n", posProp)
	fmt.Printf("%.6f\n", negProp)
	fmt.Printf("%.6f\n", zeroProp)
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
