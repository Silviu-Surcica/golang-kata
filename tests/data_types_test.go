package lessons_test

import (
   "math"
   "math/rand"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang-kata/lessons"
)

type testRecord struct {
  integers []int
  outlier  int
}

var _ = Describe("Basic tests", func() {
  It("Odd at the back", func() {
    Expect(lessons.FindOutlier([]int{2, 6, 8, 10, 3})).To(Equal(3))
  })
  It("Odd in the middle", func() {
    Expect(lessons.FindOutlier([]int{2, 6, 8, 200, 700, 1, 84, 10, 4})).To(Equal(1))
  })
  It("Odd in the front", func() {
    Expect(lessons.FindOutlier([]int{17, 6, 8, 10, 6, 12, 24, 36})).To(Equal(17))
  })
  It("Even in the front", func() {
    Expect(lessons.FindOutlier([]int{2, 1, 7, 17, 19, 211, 7})).To(Equal(2))
  })
  It("Even in the middle", func() {
    Expect(lessons.FindOutlier([]int{1, 1, 1, 1, 1, 44, 7, 7, 7, 7, 7, 7, 7, 7})).To(Equal(44))
  })
  It("Even at the end", func() {
    Expect(lessons.FindOutlier([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 35, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7, 7, 7, 7, 1000})).To(Equal(1000))
  })
  It("Odd at the back, negative", func() {
    Expect(lessons.FindOutlier([]int{2, -6, 8, -10, -3})).To(Equal(-3))
  })
  It("Odd in the middle, negative", func() {
    Expect(lessons.FindOutlier([]int{2, 6, 8, 2, -66, 34, -35, 66, 700, 1002, -84, 10, 4})).To(Equal(-35))
  })
  It("Odd in the front, negative", func() {
    Expect(lessons.FindOutlier([]int{-1 * math.MaxInt32, -18, 6, -8, -10, 6, 12, -24, 36})).To(Equal(-1 * math.MaxInt32))
  })
  It("Even in the front, negative", func() {
    Expect(lessons.FindOutlier([]int{-20, 1, 7, 17, 19, 211, 7})).To(Equal(-20))
  })
  It("Even in the middle, negative", func() {
    Expect(lessons.FindOutlier([]int{1, 1, -1, 1, 1, -44, 7, 7, 7, 7, 7, 7, 7, 7})).To(Equal(-44))
  })
  It("Odd answer, zeroes", func() {
    Expect(lessons.FindOutlier([]int{1, 0, 0})).To(Equal(1))
  })
  It("Even in the middle, zero", func() {
    Expect(lessons.FindOutlier([]int{3, 7, -99, 81, 90211, 0, 7})).To(Equal(0))
  })
})

func randOdd(n int) int {
  return rand.Intn(n/2-1)*2 + 1
}

func randEven(n int) int {
  return rand.Intn(n/2) * 2
}

func randomTest(size, maxVal int) {
  posA := rand.Intn(size)
  posB := rand.Intn(size)
  odds := make([]int, size)
  evens := make([]int, size)
  for i := 0; i < size; i++ {
    odds[i] = randOdd(maxVal)
    evens[i] = randEven(maxVal)
  }
  oddVal := randOdd(maxVal)
  evenVal := randEven(maxVal)
  odds[posA] = evenVal
  evens[posB] = oddVal
  Expect(lessons.FindOutlier(odds)).To(Equal(evenVal))
  Expect(lessons.FindOutlier(evens)).To(Equal(oddVal))
}

var _ = Describe("Random tests", func() {
  It("Small array test", func() {
    randomTest(400, 3000)
  })
  It("Medium array test", func() {
    randomTest(1000, 30000)
  })
  It("Large array test", func() {
    randomTest(100000, 100000)
  })
})
