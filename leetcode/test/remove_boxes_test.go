package main_test

import (
	solution "oj-go/leetcode/solutions/remove-boxes/1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Leetcode/Test/RemoveBoxes", func() {
	It("case 1", func() {
		Expect(solution.Run([]int{1, 3, 2, 2, 2, 3, 4, 3, 1})).To(Equal(23))
	})
	// It("case 2", func() {
	// 	Expect(solution.Run([]int{3, 8, 8, 5, 5, 3, 9, 2, 4, 4, 6, 5, 8, 4, 8, 6, 9, 6, 2, 8, 6, 4, 1, 9, 5, 3, 10, 5, 3, 3, 9, 8, 8, 6, 5, 3, 7, 4, 9, 6, 3, 9, 4, 3, 5, 10, 7, 6, 10, 7})).To(Equal(23))
	// })
})
