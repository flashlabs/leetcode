package main

import (
	"flag"
	"fmt"

	"github.com/flashlabs/leetcode/addtwonumbers"
	"github.com/flashlabs/leetcode/twosum"
)

func main() {
	fn := getInput()

	switch fn {
	case "two-sum":
		twosum.Run()
	case "add-two-numbers":
		addtwonumbers.Run()
	default:
		fmt.Println("Please provide function name. Possible options: two-sum, add-two-numbers")
	}
}

func getInput() string {
	fn := flag.String("fn", "not_existing_function_name", "Run given function")

	flag.Parse()

	return *fn
}
