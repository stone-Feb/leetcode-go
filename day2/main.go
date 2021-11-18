package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println()
	works := []string{"a","ab","abc","d","cd","bcd","abcd"}
	fmt.Println(maxProduct(works))
}

//给定一个字符串数组words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。
//假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
func maxProduct(words []string) int {
	var p int

	hasSomeChar := func(a , b string) bool {
		for _ , x := range a {
			if strings.Count(b, string(x)) != 0 {
				return true
			}
		}
		return false
	}

	for i := range words {
		//words[i] 和 words[j]两两相比，j := i + 1 保证只比较一次
		for j := i + 1; j < len(words); j++ {
			//字符串是字节的定长数组
			poi := hasSomeChar(words[i], words[j])
			if poi == false {
				tem := len(words[i]) * len(words[j])
				if p < tem {
					p = tem
				}
			}
		}
	}
	return p
}

//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
func singleNumber(nums []int) int {
	hashTable := map[int]int{}
	single := 0
	for _, item := range nums {
		hashTable[item] += 1
	}
	for k, v := range hashTable {
		if v == 1 {
			single = k
		}
	}
	return single
}
