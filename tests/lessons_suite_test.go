package lessons_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"golang-kata/lessons"
)

var _ = Describe("Basic tests", func() {
  It("Test power 1", func() {
    Expect(lessons.Pow(2, 1)).To(Equal(2.0))
  })
  It("Test power 2", func() {
    Expect(lessons.Pow(2, 2)).To(Equal(4.0))
  })
  It("Test negative power -1", func() {
    Expect(lessons.Pow(2, -1)).To(Equal(0.5))
  })
  It("Test negative power -2", func() {
    Expect(lessons.Pow(2, -2)).To(Equal(0.25))
  })
  It("Test negative base power 1", func() {
    Expect(lessons.Pow(-2, 1)).To(Equal(-2.0))
  })
  It("Test negative base power 2", func() {
    Expect(lessons.Pow(-2, 2)).To(Equal(4.0))
  })
  It("Test 0 at a power", func() {
    Expect(lessons.Pow(0, 2)).To(Equal(0.0))
  })
  It("Test negative at power 0", func() {
    Expect(lessons.Pow(-2, 0)).To(Equal(1.0))
  })
  It("Test positive at power 0", func() {
    Expect(lessons.Pow(-2, 0)).To(Equal(1.0))
  })
})

func TestLessons(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lessons Suite")
}
