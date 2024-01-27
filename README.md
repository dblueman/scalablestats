# ScalableStats
ScalableStats is a library to allow collection and processing of metrics with O(1) constant time and space guarantees; this makes it suitable for frequently-sampled metrics such as temperature.

Guarantees are provided by forming a histogram based on *apriori* knowledge of the data, for example sampled temperatures will lie in the range of 10 to 80'C, quantized to 2'C, ie 35 bins.

Arbitrary percentiles can be extracted from the resultant histogram.

## Example
    h := NewLinearHistogram(10, 80, 35)

    for i := 0; i < 100000; i++ {
        h.Store(rand.Float32() * 80)
    }

    fmt.Printf("P95=%v\n", h.Percentile(95))
    h.Fprint(os.Stdout)
    h.CLear()
