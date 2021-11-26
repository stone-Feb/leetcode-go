package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(countBits(5))
}

//给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
func countBits(n int) []int {
	count := make([]int , n+1)
	//Brian Kernighan 算法
	for i := range count {
		count[i] = func(x int) (cnt int) {
			for ; x > 0 ; x &= x-1 {
				cnt++
			}
			return
		}(i)
	}


	//系统函数库方法
	//item := ""
	//for i := 0; i <= n; i++ {
	//	item = strconv.FormatInt(int64(i), 2)
	//	count[i] = strings.Count(item,"1")
	//}
	return count
}

//二进制加法
//给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
//输入为 非空 字符串且只包含数字 1 和 0。
func addBinary(a, b string) string {
	//模拟逢二进一
	answer, carry := "", 0
	LenA, LenB := len(a), len(b)
	n := LenA
	if LenB > LenA {
		n = LenB
	}
	for i := 0; i < n; i++ {
		if i < LenA {
			carry += int(a[LenA-i-1] - '0')
		}
		if i < LenB {
			carry += int(b[LenB-i-1] - '0')
		}
		//结果在前，不用反转字符串
		answer = strconv.Itoa(carry%2) + answer
		//二进制求进位
		carry /= 2
	}
	//最后一个进位的计算
	if carry > 0 {
		answer = "1" + answer
	}
	return answer
}

//整数除法
//给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%'。
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
