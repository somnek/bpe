package main

import "fmt"

func main() {
	// 'ids' in karpathy's video
	// normalize to 0 - 256 (convert bytes to runes)
	// convert string to runes right away won't works because the values goes beyond 256
	bytes := []byte(WIKI)
	runes := bytesToRunes(bytes)
	stats := getStats(runes)
	fmt.Println(stats)

}

type Pair struct {
	a, b rune
}

func getStats(runes []rune) map[Pair]int {
	m := make(map[Pair]int)
	for i := 0; i < len(runes)-1; i++ {
		pair := Pair{runes[i], runes[i+1]}
		m[pair]++
	}
	return m
}

func bytesToRunes(bytes []byte) []rune {
	var runeList []rune
	for _, b := range bytes {
		runeList = append(runeList, rune(b))
	}
	return runeList
}

func runesToBytes(runes []rune) []byte {
	var byteList []byte
	for _, r := range runes {
		byteList = append(byteList, byte(r))
	}
	return byteList
}
