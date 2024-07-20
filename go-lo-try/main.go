package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	items := lo.Uniq([]string{"A", "A", "B", "A", "C"})
	fmt.Printf("%+v\n", items)

	values := []string{"A", "A", "B", "A", "C"}
	filtered := lo.Filter(values, func(x string, index int) bool {
		return x != "B"
	})
	fmt.Printf("%+v\n", filtered)
}
