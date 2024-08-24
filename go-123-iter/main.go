package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
	"time"
)

func main() {
	// usecase 1
	items := []string{"a", "b", "c"}
	itemsIter := slices.Values(items)
	fmt.Printf("%+v\n", itemsIter)

	for v := range convert(itemsIter) {
		fmt.Printf("v %s\n", v)
		time.Sleep(1 * time.Second)
	}

	// usecase 2
    for text := range splitext("aaa,aab,bbccc") {
		fmt.Println(text)
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

func splitext(text string) iter.Seq[string] {
	// yield というキーワードがある訳ではない
	return func(yield func(string) bool) {
		items := strings.Split(text, ",")
		// do something here
		for _, item := range items {
			yield(item)
		}
	}
}
