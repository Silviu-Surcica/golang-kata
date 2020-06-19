package lessons_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang-kata/lessons"
)

var _ = Describe("DataTypes", func() {
    It("Test abc", func() {
    Expect(lessons.VowelOrConsonant("abc")).To(Equal(map[string][]string{
                                                         "Vowels": {"a"},
                                                         "Consonants": {"b", "c"},
                                                         },
    ))
  })
})
