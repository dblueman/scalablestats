package scalablestats

import (
   "fmt"
   "os"
)

type Histogram struct {
   Thresholds []float32
   Counts     []int
}

func NewLinearHistogram(min, max, bins int) *Histogram {
   h := Histogram{}

   h.Counts = make([]int, bins)

   for b := 0; b < bins; b++ {
      h.Thresholds = append(h.Thresholds, float32(min) + (float32((max - min) * (b+1)) / float32(bins)))
   }

   return &h
}

func index(val float32, thresholds []float32) int {
   for i, t := range thresholds {
      if val < t {
         return i
      }
   }

   return len(thresholds)-1
}

func (h *Histogram) Store(val float32) {
   i := index(val, h.Thresholds)
   h.Counts[i]++
}

func (h *Histogram) Percentile(threshold int) float32 {
   counts := 0

   for _, count := range(h.Counts) {
      counts += count
   }

   limit := counts * 100 / threshold
   total := 0

   for i := 0; i < len(h.Counts); i++ {
      total += h.Counts[i]

      if i > limit {
         return h.Thresholds[i]
      }
   }

   return h.Thresholds[len(h.Thresholds)-1]
}

func (h *Histogram) Clear() {
   for i := range h.Counts {
      h.Counts[i] = 0
   }
}

func (h *Histogram) Fprint(f *os.File) {
   fmt.Fprintf(f, "thresholds %+v\n", h.Thresholds)
   fmt.Fprintf(f, "counts %+v\n", h.Counts)
}
