package main

import "fmt"

///////////////////////
// SOLUTION 1
///////////////////////

// func twoSum(nums []int, target int) [2]int {

// 	for index1, item1 := range nums {
// 		for index2, item2 := range nums {
// 			if index1 >= index2 {
// 				continue
// 			}

// 			if item1+item2 == target {
// 				return [2]int{index1, index2}
// 			}
// 		}
// 	}

// 	return [2]int{-1, -1}
// }

///////////////////////
// SOLUTION 2
///////////////////////

// func twoSum(nums []int, target int) []int {

// 	hashmap := make(map[int]int)

// 	for index, item := range nums {
// 		hashmap[item] = index
// 	}

// 	for index, item := range nums {
// 		diff := target - item
// 		if _, ok := hashmap[diff]; ok && index != hashmap[diff] {
// 			return []int{index, hashmap[diff]}
// 		}
// 	}

// 	return []int{-1, -1}
// }

///////////////////////
// SOLUTION 3
///////////////////////

func twoSum(nums []int, target int) []int {
	seenNums := make(map[int]int)
	for index, thisNum := range nums {
		if seenIndex, ok := seenNums[target-thisNum]; ok {
			return []int{seenIndex, index}
		}
		seenNums[thisNum] = index
	}
	return []int{-1, -1}
}

///////////////////////
// DRIVER
///////////////////////

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}
