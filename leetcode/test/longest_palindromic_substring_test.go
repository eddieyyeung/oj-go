package main_test

import (
	solution "oj-go/leetcode/solutions/longest-palindromic-substring/1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Leetcode/Test/LongestPalindromicSubstring", func() {
	It("case 1", func() {
		Expect(solution.Run("babad")).To(Equal("aba"))
	})
})
