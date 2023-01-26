package main

import (
	"fmt"
)

func main() {
	sequence := "ATGCATTGC"
	subSequence := "CTG"

	seqLen := len(sequence)
	subSeqLen := len(subSequence)

	var truncatedSequence string
	outMap := make(map[int]float64)
	var ch = make(chan res)

	for i := 0; i < seqLen; i++ {
		if i+subSeqLen > seqLen {
			truncatedSequence = sequence[i:seqLen]
		} else {
			truncatedSequence = sequence[i : i+subSeqLen]
		}
		go comp(truncatedSequence, subSequence, i, ch)
	}

	for i := 0; i < seqLen; i++ {
		resp := <-ch
		outMap[resp.key] = resp.value
	}

	for key, value := range outMap {
		fmt.Println(key, value)
	}
}

type res struct {
	key   int
	value float64
}

func comp(seq string, subSeq string, key int, c chan res) {
	seqLen := len(seq)
	subSeqLen := len(subSeq)

	var count int
	for i := 0; i < seqLen; i++ {
		if seq[i] == subSeq[i] {
			count++
		}
	}

	c <- res{key, float64(count) / float64(subSeqLen)}
}
