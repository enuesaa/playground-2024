package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	names := lo.Uniq([]string{"Samuel", "John", "Samuel"})
	fmt.Println(names)

	// 三項演算子もどき
	res := lo.Ternary(true, "a", "b")
	fmt.Println(res)
}
