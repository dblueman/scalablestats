package scalablestats

import (
   "fmt"
   "math/rand"
   "testing"
)

func Test(t *testing.T) {
   h := NewLinearHistogram(0, 1000, 20)

   for i := 0; i < 1000000; i++ {
      h.Store(float32(rand.NormFloat64() * 100. + 500))
   }

   if h.Percentile(50) != 500 {
      t.Error("incorrect median")
   }

   h.Clear()
}
