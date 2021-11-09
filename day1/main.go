package main

import (
	"fmt"
	"math"
)

func main() {
	//devide()
	var a, b int32
	a = 10
	b = 2
	fmt.Println("result :", divide(a, b))
	//fmt.Println(5<<4)
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

	// 时间复杂度为log(n)
	//for a >= b {
	//	value , k := b , 1
	//	for a >= value + value {
	//		value += value
	//		k += k
	//	}
	//	a -= value
	//	count += k
	//}

	//位运算 时间复杂度为 O(n)
	for i := 31; i >= 0; i-- {
		if b<<i > 0 && a-(b<<i) >= 0 {
			a -= b << i
			count += 1 << i
		}
	}

	if sign == -1 {
		count = -count
	}
	return count
}
