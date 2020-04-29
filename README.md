# go-calculator

<a href="https://github.com/greenpau/go-calculator/actions/" target="_blank"><img src="https://github.com/greenpau/go-calculator/workflows/build/badge.svg?branch=master"></a>
<a href="https://pkg.go.dev/github.com/greenpau/go-calculator" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"></a>
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)

Go package to calculate total, mean (average), median, mode, range, etc.

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
c.Total()
c.StandardDeviation()
c.Variance()
c.Range()
c.Max()
c.Min()
c.Median(true)
c.Median(false)
c.Mean()
c.Modes()
```

Alternatively, simply perform all calculations:

```
calc.RunAll()
```

Get the result of the calculations.

```golang
fmt.Fprintf(os.Stdout, "Total: %.2f\n", calc.Register.Total)
```
