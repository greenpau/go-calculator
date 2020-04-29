# go-calculator

<a href="https://github.com/greenpau/go-calculator/actions/" target="_blank"><img src="https://github.com/greenpau/go-calculator/workflows/build/badge.svg?branch=master"></a>
<a href="https://pkg.go.dev/github.com/greenpau/go-calculator" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"></a>
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)

Go package to calculate total, mean (average), median, sorted median,
variance, range, min, max, modes, standard deviation, etc.

## Getting Started

First, add the library to :

```golang
import (
    "github.com/greenpau/go-calculator"
)
```

Next, initialize calculator in any of the following ways:

```golang
calc, _ := calculator.NewString("1, 2, 3, 4.5, 5.4, 6, 7")

if calc == nil {
    log.Fatal("failed to initialize calculator")
}

arr := []uint64{1, 2, 3, 4, 5, 6, 7}
calc, _ := calculator.NewUint64(arr)

arr := []int64{1, 2, 3, 4, 5, 6, 7}
calc, _ := calculator.NewInt64(arr)

arr := []int{1, 2, 3, 4, 5, 6, 7}
calc, _ := calculator.NewInt(arr)

arr := []float64{1, 2, 3, 4, 5, 6, 7}
calc, _ := calculator.NewFloat64(arr)
```

Next, calculate total, mean (average), median, mode, range
using the `calc` instance:

```golang
cacl.Total()
calc.StandardDeviation()
calc.Variance()
calc.Range()
calc.Max()
calc.Min()
calc.MaxWithIndices()
calc.MinWithIndices()
calc.Median(true)
calc.Median(false)
calc.Mean()
calc.Modes()
```

Alternatively, simply perform all calculations:

```
calc.RunAll()
```

Get the result of the calculations.

```golang
fmt.Fprintf(os.Stdout, "Total: %.2f\n", calc.Register.Total)
```

The `Print()` function helps visualize the data:

```
Data: [554801 61602 763767 277578 167822 617342
  847743 801166 894535 254904 354657 389371
  543430]
Total: 6528718.000000
Mean: 502209.076923
Median: 732542.500000
Sorted Median: 466400.500000
Max: 894535.000000
Max Indices: [8]
Min: 61602.000000
Min Indices: [1]
Variance: 69299542922.378693
Standard Deviation: 0.000000
Modes: []
Mode Repeat Count: 0
```
