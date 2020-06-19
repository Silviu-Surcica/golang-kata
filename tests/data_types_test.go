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
  It("Test aeo", func() {
    Expect(lessons.VowelOrConsonant("aeo")).To(Equal(map[string][]string{
                                                         "Vowels": {"a", "e", "o"},
                                                         "Consonants": {},
                                                         },
    ))
  })
  It("Test aaaa", func() {
    Expect(lessons.VowelOrConsonant("aaaa")).To(Equal(map[string][]string{
                                                         "Vowels": {"a", "a", "a", "a"},
                                                         "Consonants": {},
                                                         },
    ))
  })
  It("Test a", func() {
    Expect(lessons.VowelOrConsonant("a")).To(Equal(map[string][]string{
                                                         "Vowels": {"a"},
                                                         "Consonants": {},
                                                         },
    ))
  })
  It("Test aeiou", func() {
    Expect(lessons.VowelOrConsonant("aeiou")).To(Equal(map[string][]string{
                                                         "Vowels": {"a", "e", "i", "o", "u"},
                                                         "Consonants": {},
                                                         },
    ))
  })
  It("Test c", func() {
    Expect(lessons.VowelOrConsonant("c")).To(Equal(map[string][]string{
                                                         "Vowels": {},
                                                         "Consonants": {"c"},
                                                         },
    ))
  })
  It("Test empty", func() {
    Expect(lessons.VowelOrConsonant("")).To(Equal(map[string][]string{
                                                         "Vowels": {},
                                                         "Consonants": {},
                                                         },
    ))
  })
})
