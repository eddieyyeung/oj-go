package multiplesof3or5

// Multiple3And5 https://www.codewars.com/kata/514b92a657cdc65150000006/go
func Multiple3And5(number int) int {
	n := number - 1
	d3 := n / 3
	sum3 := (3 + d3*3) * d3 / 2
	d5 := n / 5
	sum5 := (5 + d5*5) * d5 / 2
	d15 := n / 15
	sum15 := (15 + d15*15) * d15 / 2
	return sum3 + sum5 - sum15
}
