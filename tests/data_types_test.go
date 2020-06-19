package lessons_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang-kata/lessons"
)

var _ = Describe("DataTypes", func() {
    It("Test 12", func() {
    Expect(lessons.DigitalRoot(12)).To(Equal(3))
  })
  It("Test 195", func() {
    Expect(lessons.DigitalRoot(195)).To(Equal(6))
  })
  It("Test 993", func() {
    Expect(lessons.DigitalRoot(993)).To(Equal(3))
  })
    It("Test 167346", func() {
        Expect(lessons.DigitalRoot(167346)).To(Equal(9))
      })
      It("Test 0", func() {
        Expect(lessons.DigitalRoot(0)).To(Equal(0))
      })
})
