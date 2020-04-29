# go-calculator

<a href="https://github.com/greenpau/go-calculator/actions/" target="_blank"><img src="https://github.com/greenpau/go-calculator/workflows/build/badge.svg?branch=master"></a>
<a href="https://pkg.go.dev/github.com/greenpau/go-calculator" target="_blank"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"></a>
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)

Calculate total, mean (average), median, mode, and range.

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
result := calc.Total()
result := calc.Mean()
result := calc.Average() // shortcut to Mean()
result := calc.Median()
result := calc.Mode()
result := calc.Range()
```

The `result` is also an object. It is capable outputing `uint64`, `int64`,
`int`, or `float64`.

```golang
fmt.Fprintf(os.Stdout, "Total (string): %d\n", calc.Total().String())
fmt.Fprintf(os.Stdout, "Total (uint64): %d\n", calc.Total().Uint64())
fmt.Fprintf(os.Stdout, "Total (int64): %d\n", calc.Total().Int64())
fmt.Fprintf(os.Stdout, "Total (int): %d\n", calc.Total().Int())
fmt.Fprintf(os.Stdout, "Total (float64): %.2f\n", calc.Total().Float64())
```
