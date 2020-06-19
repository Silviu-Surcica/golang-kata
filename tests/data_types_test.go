package lessons_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang-kata/lessons"
)

var _ = Describe("DataTypes", func() {
    It("Test 'there was'", func() {
    Expect(lessons.WordLength("there was")).To(Equal(map[string]int{
                                                         "there": 5,
                                                         "was": 3,
                                                         },
    ))
  })
  It("Test 'I am.'", func() {
    Expect(lessons.WordLength("I am.")).To(Equal(map[string]int{
                                                         "I": 1,
                                                         "am.": 3,
                                                         },
    ))
  })
  It("Test 'Us all, are we.'", func() {
    Expect(lessons.WordLength("Us all, are we.")).To(Equal(map[string]int{
                                                         "Us": 2,
                                                         "all,": 4,
                                                         "are": 3,
                                                         "we.": 3,
                                                         },
    ))
  })
  It("Test empty", func() {
    Expect(lessons.WordLength("")).To(Equal(map[string]int{

                                                         },
    ))
  })
  It("Test signs '! , . 0 @ #'", func() {
    Expect(lessons.WordLength("! , . 0 @ #")).To(Equal(map[string]int{
                                                         "!": 1,
                                                         ",": 1,
                                                         ".": 1,
                                                         "0": 1,
                                                         "@": 1,
                                                         "#": 1,
                                                         },
    ))
  })
})
