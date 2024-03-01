package main

import (
	"fmt"
	"strings"
)

func main() {
	merges, encoded := encode(WHOLE_TEXT)
	encodedStr := string(runesToBytes(encoded))
	fmt.Printf("encoded: \n%s\n    length: %d\n", encodedStr, len(encodedStr))

	decoded := decode(merges, encoded)
	decodedStr := string(runesToBytes(decoded))
	fmt.Printf("decoded: \n%v\n", decodedStr)

	fmt.Printf("accurately decoded: %v\n    length: %d\n", decodedStr == WHOLE_TEXT, len(decodedStr))
}

type Pair struct {
	a, b rune
}

type Row struct {
	pair  Pair
	alien rune
}

// -----------------------------------------------------------------------------
//
//	COMPRESS/ ENCODE
//
// -----------------------------------------------------------------------------
func encode(text string) ([]Row, []rune) {
	// 'ids' in karpathy's video
	// normalize to 0 - 256 (bytes, which has longer length)
	// convert string to runes right away won't works because the rune values > 256 (we want to stay i 0-256)
	// thats why we convert to bytes first then runes, which keep same length as original
	bytes := []byte(text)
	runes := bytesToRunes(bytes)
	var alien int32
	iteration := VOCAB_SIZE - 256
	// use for decoding (FIFO)
	var merges []Row

	for j := 0; j < iteration; j++ {
		stats := getStats(runes)
		maxPair, maxCount := findMaxPair(stats)
		if maxCount <= 1 {
			break
		}
		alien = 256 + rune(j)
		runes = merge(runes, maxPair, alien)
		// NOTE: to convert back, turn into []rune first then turn the rune to byte
		// like so: string(runesToBytes(runes))
		merges = append([]Row{{maxPair, alien}}, merges...) // newest on top
	}
	return merges, runes
}

// -----------------------------------------------------------------------------
//
//	DECOMPRESS/ DECODE
//
// -----------------------------------------------------------------------------
func decode(merges []Row, runes []rune) []rune {
	for _, e := range merges {
		pairJoined := string(e.pair.a) + string(e.pair.b)
		runes = []rune(strings.ReplaceAll(string(runes), string(e.alien), pairJoined))
	}
	return runes
}

func merge(runes []rune, pair Pair, alien rune) []rune {
	pairJoined := string(pair.a) + string(pair.b)
	s := string(runes)
	s = strings.ReplaceAll(s, pairJoined, string(alien))
	return []rune(s)
}

func findMaxPair(stats map[Pair]int) (Pair, int) {
	var maxCount int
	var maxPair Pair
	for pair, count := range stats {
		if count > maxCount {
			maxPair = pair
			maxCount = count
		}
	}
	return maxPair, maxCount
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
