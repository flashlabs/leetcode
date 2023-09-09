package twosum

import "fmt"

func Run() {
	fmt.Println("TWO SUM")

	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	for idx, val := range nums {
		for idx2, val2 := range nums {
			if idx2 == idx {
				continue
			}
			if val+val2 == target {
				return []int{idx, idx2}
			}
		}
	}
	return nil
}
