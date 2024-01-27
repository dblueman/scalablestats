package scalablestats

import (
   "fmt"
   "os"
   "slices"
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

func (h *Histogram) Store(val float32) {
   i, _ := slices.BinarySearch(h.Thresholds, val)
   h.Counts[i]++
}

func (h *Histogram) Percentile(threshold int) float32 {
   counts := 0

   for _, count := range(h.Counts) {
      counts += count
   }

   limit := counts * threshold / 100
   total := 0

   for i := 0; i < len(h.Counts); i++ {
      total += h.Counts[i]

      if total > limit {
         return h.Thresholds[i]
      }
   }

   return h.Thresholds[len(h.Thresholds)-1]
}

func (h *Histogram) Clear() {
   clear(h.Counts)
}

func (h *Histogram) Fprint(f *os.File) {
   fmt.Fprintf(f, "thresholds %+v\n", h.Thresholds)
   fmt.Fprintf(f, "counts %+v\n", h.Counts)
}
