package main

import (
	"fmt"

	"github.com/biuboombiuboom/leetcode/problem7"
)

func main() {

	// 717
	// inputs := []int{1, 1, 1, 0}
	// fmt.Println(strconv.FormatBool(problem717.IsOneBitCharacter(inputs)))

	// 1
	// nums := []int{3, 2, 4}
	// target := 6
	// result := problem1.TwoSum(nums, target)

	// 7
	input := -123
	result := problem7.Reverse(input)
	fmt.Println(result)
}
