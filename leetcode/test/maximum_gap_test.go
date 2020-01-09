package main_test

import (
	solution "oj-go/leetcode/solutions/maximum-gap/1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Leetcode/Test/MaximumGap", func() {
	It("[3, 6, 9, 1]", func() {
		Expect(solution.Run([]int{3, 6, 9, 1})).To(Equal(3))
	})
})
