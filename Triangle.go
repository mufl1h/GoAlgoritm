package main

import(
	"fmt"
)

func main() {
	fmt.Println(tirangleRight(6))
	fmt.Println(tirangleLeft(6))
}

func tirangleRight(n int) string {
	var ret string = ""
	for i := 1; i <= n; i ++ {
		for j := 0; j < n - i; j ++ {
			ret += " "
		}

		for j := 0; j < i; j ++ {
			ret += "#"
		}
		ret += "\n"
	}
	return ret
}

func tirangleLeft(n int) string {
	var ret string = ""
	for i := 1; i <= n; i ++ {
		for j := 0; j < i; j ++ {
			ret += "#"
		}
		ret += "\n"
	}
	return ret
}