package codewars_test

import (
	. "oj-go/codewars/solution"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(n int, exp int) {
	var ans = ARuleOfDivisibilityBy13(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("should handle basic cases", func() {
		testequal(8529, 79)
		testequal(85299258, 31)
		testequal(5634, 57)
	})
})
