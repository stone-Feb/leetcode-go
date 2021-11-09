package main

import (
	"fmt"
	"math"
)

func main() {
	//devide()
	//var a, b int32
	a := 10000
	//b := 2
	//fmt.Println(divide(a, b))
	fmt.Println(a>>1)
}

//给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%'。
//
//假设我们的环境只能存储 32 位有符号整数
func divide(a, b int32) int {
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt
	}
	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	count := 0
	// 时间复杂度为O(n)
	//for a >= b {
	//	a -= b
	//	count++
	//}

	if sign == -1 {
		count = -count
	}
	return count
}
