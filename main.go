package main

import (
	"fmt"
	"strings"
)

type Pair struct {
	a, b byte
}

func main() {
	// bytes := []byte(WIKI)
	sWiki := WIKI
	out := 90

	for j := 0; j < MAX_ITER; j++ {
		m := make(map[Pair]int)
		for i := 0; i < len(sWiki)-1; i++ {
			pair := Pair{sWiki[i], sWiki[i+1]}
			m[pair]++
		}
		// find max pair
		maxPair, maxCount := findMaxPair(m)
		if maxCount <= 1 {
			break
		}
		joinedPair := string(maxPair.a) + string(maxPair.b)
		sWiki = strings.ReplaceAll(sWiki, joinedPair, string(rune(out-j)))
	}
	fmt.Println("final: ", sWiki)
}

func findMaxPair(m map[Pair]int) (Pair, int) {
	var maxCount int
	var maxPair Pair
	for p, count := range m {
		if count >= maxCount {
			maxCount = count
			maxPair = p
		}
	}
	return maxPair, maxCount
}
