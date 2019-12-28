package codewars_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "oj-go/codewars/solution"
)

func dotest(n int64, exp int64) {
	var ans = RotateForAMax(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Codewars/Test/RotateForAMax", func() {

	It("should handle basic cases", func() {
		dotest(38458215, 85821534)
		dotest(195881031, 988103115)
		dotest(896219342, 962193428)

	})
})
