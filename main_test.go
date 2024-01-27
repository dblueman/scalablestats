package scalablestats

import (
   "fmt"
   "math"
   "math/rand"
   "os"
   "testing"
)

func Test(t *testing.T) {
   h := NewLinearHistogram(0, 1000, 20)

   for i := 0; i < 100000; i++ {
      h.Store(float32(math.Abs(rand.NormFloat64()) * 200.))
   }

   fmt.Printf("P95=%v\n", h.Percentile(95))
   h.Fprint(os.Stdout)
}
