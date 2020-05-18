package main

import (
	seq "./sequence"
	"fmt"
)

const nodeId = 125

func log(s seq.Sequence) {
	fmt.Printf("IDs: %d\n", s.CurrVal())
}

func main() {
	seq1 := seq.Snowflake{Node: seq.Node{Id: nodeId}}

	for i := 0; i <= 100; i++ {

		//time.Sleep(500)
		seq1.NextVal()
		log(&seq1)
	}

	//seq2 := new(seq.Snowflake)
	//seq2.Node = seq.Node{Id: NODE}
	//log(seq2)
}
