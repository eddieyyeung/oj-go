// Package multiply_strings ...
// https://leetcode-cn.com/problems/multiply-strings/
package multiply_strings

import (
	"math"
)

func maxf(a ...int) int {
	max := math.MinInt64
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func add(num1 string, num2 string) string {
	carry := 0
	l1 := len(num1) - 1
	l2 := len(num2) - 1
	r := make([]rune, maxf(len(num1), len(num2))+1)
	l := len(r) - 1
	for l1 >= 0 && l2 >= 0 {
		n1 := int(num1[l1] - '0')
		n2 := int(num2[l2] - '0')
		sum := n1 + n2 + carry
		carry = sum / 10
		r[l] = rune(sum%10 + '0')
		l--
		l1--
		l2--
	}
	for l1 >= 0 {
		n1 := int(num1[l1] - '0')
		sum := n1 + carry
		carry = sum / 10
		r[l] = rune(sum%10 + '0')
		l--
		l1--
	}
	for l2 >= 0 {
		n2 := int(num2[l2] - '0')
		sum := n2 + carry
		carry = sum / 10
		r[l] = rune(sum%10 + '0')
		l--
		l2--
	}
	if carry != 0 {
		r[l] = rune(carry + '0')
		l--
	}
	return string(r[l+1:])
}

func mulOne(one int, num string) string {
	l := len(num) - 1
	carry := 0
	r := make([]rune, len(num)+1)
	u := len(r) - 1
	for l >= 0 {
		n := int(num[l] - '0')
		mul := n*one + carry
		carry = mul / 10
		r[u] = rune(mul%10 + '0')
		u--
		l--
	}
	if carry != 0 {
		r[u] = rune(carry + '0')
		u--
	}
	return string(r[u+1:])
}

func multiply(num1 string, num2 string) string {
	if len(num1) == 1 && num1[0] == '0' {
		return "0"
	}
	if len(num2) == 1 && num2[0] == '0' {
		return "0"
	}
	r := "0"
	ten := ""
	l := len(num1) - 1
	for l >= 0 {
		n1 := int(num1[l] - '0')
		mul := mulOne(n1, num2) + ten
		r = add(r, mul)
		ten += "0"
		l--
	}
	return r
}
