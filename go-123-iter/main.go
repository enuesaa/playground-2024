package main

import (
	"fmt"
	"iter"
	"slices"
	"time"
)

func main() {
	items := []string{"a", "b", "c"}
	itemsIter := slices.Values(items)
	fmt.Printf("%+v\n", itemsIter)

	for v := range convert(itemsIter) {
		fmt.Printf("v %s\n", v)
		time.Sleep(1 * time.Second)
	}
}

func convert(seq iter.Seq[string]) iter.Seq[string] {
	// see https://pkg.go.dev/iter#Pull
	next, stop := iter.Pull(seq)
	value, _ := next()
	fmt.Printf("value %s\n", value)
	stop()

	fmt.Printf("seq %+v\n", seq)
	return seq
}
