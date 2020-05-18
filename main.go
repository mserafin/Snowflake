package main

import (
    seq "./sequence"
    "fmt"
)

const nodeId = 125

func log(seq seq.Sequence) {
    fmt.Printf("PK: %d\n", seq.CurrVal())
}

func main() {
    seq := seq.Snowflake{Node: seq.Node{Id: nodeId}}

    for i := 0; i <= 100; i++ {
        seq.NextVal()
        log(&seq)
    }
}
