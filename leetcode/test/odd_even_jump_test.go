package main_test

import (
	solution "oj-go/leetcode/solutions/odd-even-jump/4"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Leetcode/Test/OddEvenJump", func() {
	It("[10, 13, 12, 14, 15]", func() {
		Expect(solution.Run([]int{10, 13, 12, 14, 15})).To(Equal(2))
	})
	It("[2, 3, 1, 1, 4]", func() {
		Expect(solution.Run([]int{2, 3, 1, 1, 4})).To(Equal(3))
	})
	It("[5, 1, 3, 4, 2]", func() {
		Expect(solution.Run([]int{5, 1, 3, 4, 2})).To(Equal(3))
	})
	It("[1, 2, 3, 2, 1, 4, 4, 5]", func() {
		Expect(solution.Run([]int{1, 2, 3, 2, 1, 4, 4, 5})).To(Equal(6))
	})
	It("[5, 1, 3, 4, 2]", func() {
		Expect(solution.Run([]int{5, 1, 3, 4, 2})).To(Equal(3))
	})
	It("[2, 3, 1, 1, 4, 2, 3, 1, 1, 4]", func() {
		Expect(solution.Run([]int{2, 3, 1, 1, 4, 2, 3, 1, 1, 4})).To(Equal(5))
	})
	It("[82, 66, 0, 49]", func() {
		Expect(solution.Run([]int{82, 66, 0, 49})).To(Equal(2))
	})
})
