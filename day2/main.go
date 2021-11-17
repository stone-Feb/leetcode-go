package main

import "fmt"

func main()  {
	nums := []int{0,1,0,1,0,1,100}
	fmt.Println(singleNumber(nums))
}



//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
func singleNumber(nums []int) int {
	hashTable := map[int]int{}
	single := 0
	for _,item := range nums {
		hashTable[item] += 1
	}
	for k,v := range hashTable {
		if v == 1 {
			single = k
		}
	}
	return single
}
