package main

import (
	"fmt"
	"strings"
)

type Pair struct {
	a, b byte
}

func main() {
	bytes := []byte(WIKI) // a.k.a 'ids' in karpathy's video
	merges := []Pair{}
	alien := 90
	// encode
	for j := 0; j < MAX_ITER; j++ {
		fmt.Println(string(bytes))
		stats := getStats(bytes)
		maxPair, maxCount := findMaxPair(stats)
		if maxCount <= 1 {
			break
		}
		bytes = merge(bytes, maxPair, rune(alien-j))
		merges = append(merges, maxPair)
	}

}

func merge(bytes []byte, pair Pair, alien rune) []byte {
	pairS := string(pair.a) + string(pair.b)
	bytes = []byte(strings.ReplaceAll(string(bytes), pairS, string(alien)))
	fmt.Printf(" %s -> %s\n", pairS, string(alien))
	return bytes
}

func getStats(bytes []byte) map[Pair]int {
	m := make(map[Pair]int)
	for i := 0; i < len(bytes)-1; i++ {
		pair := Pair{
			a: bytes[i],
			b: bytes[i+1],
		}
		m[pair]++
	}
	return m
}

func findMaxPair(m map[Pair]int) (Pair, int) {
	var maxPair Pair
	var maxCount int
	for pair, count := range m {
		if count > maxCount {
			maxCount = count
			maxPair = pair
		}
	}
	return maxPair, maxCount
}
