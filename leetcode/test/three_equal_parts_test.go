package main_test

import (
	solution "oj-go/leetcode/solutions/three-equal-parts/1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Leetcode/Test/ThreeEqualParts", func() {
	It("[0, 0, 0, 0, 0]", func() {
		Expect(solution.Run([]int{0, 0, 0, 0, 0})).To(Equal([]int{0, 2}))
	})
	It("[1, 0, 1, 0, 1]", func() {
		Expect(solution.Run([]int{1, 0, 1, 0, 1})).To(Equal([]int{0, 3}))
	})
	It("[0, 1, 0, 1, 1]", func() {
		Expect(solution.Run([]int{0, 1, 0, 1, 1})).To(Equal([]int{1, 4}))
	})
	It("[1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0]", func() {
		Expect(solution.Run([]int{1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0})).To(Equal([]int{15, 32}))
	})
})
