package codewars_test

import (
	. "oj-go/codewars/solution"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("just o", func() {
		Expect(MakeTheDeadfishSwim("ooo")).To(Equal([]int{0, 0, 0}))
	})
	It("o&i", func() {
		Expect(MakeTheDeadfishSwim("ioioio")).To(Equal([]int{1, 2, 3}))
	})
	It("o&i&d", func() {
		Expect(MakeTheDeadfishSwim("idoiido")).To(Equal([]int{0, 1}))
	})
	It("o&i&d&s", func() {
		Expect(MakeTheDeadfishSwim("isoisoiso")).To(Equal([]int{1, 4, 25}))
	})
	It("codewars", func() {
		Expect(MakeTheDeadfishSwim("codewars")).To(Equal([]int{0}))
	})
})
