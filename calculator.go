// MIT License
// Copyright 2020 Paul Greenberg (greenpau@outlook.com)

package calculator

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Cell provides the facility to perform mathematical operations.
type Cell struct {
	data                        []float64
	sortedData                  []float64
	length                      int
	even                        bool
	middleIndex                 int
	err                         error
	calculatedTotal             bool
	calculatedMean              bool
	calculatedMedian            bool
	calculatedRange             bool
	calculatedVariance          bool
	calculatedStandardDeviation bool
	calculatedSortedMedian      bool
	calculatedMax               bool
	calculatedMaxWithIndices    bool
	calculatedMin               bool
	calculatedMinWithIndices    bool
	calculatedModes             bool
	Register                    Register
}

// Register provides the means to store calculation results.
type Register struct {
	Total             float64
	Mean              float64
	Median            float64
	Range             float64
	Variance          float64
	StandardDeviation float64
	SortedMedian      float64
	MaxIndices        []int
	MaxValue          float64
	MinIndices        []int
	MinValue          float64
	ModeRepeatCount   int
	Modes             []float64
}

// Length returns Cell data element size.
func (c *Cell) Length() int {
	return c.length
}

// Even return true if the number of element in the Cell is even.
func (c *Cell) Even() bool {
	return c.even
}

func (c *Cell) validate() bool {
	if c.err != nil {
		return false
	}
	return true
}

// Total returns the total value of values in Cell.
func (c *Cell) Total() *Cell {
	if c.calculatedTotal || c.err != nil {
		return c
	}
	c.Register.Total = 0
	for _, v := range c.data {
		c.Register.Total += v
	}
	c.calculatedTotal = true
	return c
}

// Mean returns the mean (average) value of values in Cell.
func (c *Cell) Mean() *Cell {
	if c.calculatedMean || c.err != nil {
		return c
	}
	if !c.calculatedTotal {
		c.Total()
		if c.err != nil {
			return c
		}
	}
	c.Register.Mean = c.Register.Total / float64(c.length)
	c.calculatedMean = true
	return c
}

// Variance calculates variance of the values in Cell.
func (c *Cell) Variance() *Cell {
	if c.calculatedVariance || c.err != nil {
		return c
	}
	if !c.calculatedMean {
		c.Mean()
		if c.err != nil {
			return c
		}
	}
	for _, i := range c.data {
		r := c.Register.Mean - i
		c.Register.Variance += r * r
	}
	c.Register.Variance /= float64(c.length)
	c.calculatedVariance = true
	return c
}

// Range calculates range of the values in Cell.
func (c *Cell) Range() *Cell {
	if c.calculatedRange || c.err != nil {
		return c
	}

	if !c.calculatedMax {
		c.Max()
		if c.err != nil {
			return c
		}
	}

	if !c.calculatedMin {
		c.Min()
		if c.err != nil {
			return c
		}
	}

	c.Register.Range = c.Register.MaxValue - c.Register.MinValue
	c.calculatedRange = true
	return c
}

// Max calculates the biggest value in Cell.
func (c *Cell) Max() *Cell {
	if c.calculatedMax || c.err != nil {
		return c
	}
	c.Register.MaxValue = c.sortedData[c.length-1]
	c.calculatedMax = true
	return c
}

// MaxWithIndices calculates the biggest value and associated indices in Cell.
func (c *Cell) MaxWithIndices() *Cell {
	if c.calculatedMaxWithIndices || c.err != nil {
		return c
	}

	if !c.calculatedMax {
		c.Max()
		if c.err != nil {
			return c
		}
	}
	c.Register.MaxIndices = []int{}
	if c.length == 1 {
		c.Register.MaxIndices = append(c.Register.MaxIndices, 0)
		c.calculatedMaxWithIndices = true
		return c
	}
	for i, v := range c.data {
		if v == c.Register.MaxValue {
			c.Register.MaxIndices = append(c.Register.MaxIndices, i)
		}
	}
	c.calculatedMaxWithIndices = true
	return c
}

// Min calculates the smallest value in Cell.
func (c *Cell) Min() *Cell {
	if c.calculatedMin || c.err != nil {
		return c
	}
	c.Register.MinValue = c.sortedData[0]
	c.calculatedMax = true
	return c
}

// MinWithIndices calculates the smallest value and associated indices in Cell.
func (c *Cell) MinWithIndices() *Cell {
	if c.calculatedMinWithIndices || c.err != nil {
		return c
	}

	if !c.calculatedMin {
		c.Min()
		if c.err != nil {
			return c
		}
	}
	c.Register.MinIndices = []int{}
	if c.length == 1 {
		c.Register.MinIndices = append(c.Register.MinIndices, 0)
		c.calculatedMinWithIndices = true
		return c
	}
	for i, v := range c.data {
		if v == c.Register.MinValue {
			c.Register.MinIndices = append(c.Register.MinIndices, i)
		}
	}
	c.calculatedMinWithIndices = true
	return c
}

// Modes calculates the values appearing most often in Cell.
func (c *Cell) Modes() *Cell {
	if c.calculatedModes || c.err != nil {
		return c
	}
	occurences := make(map[float64]int)
	c.Register.Modes = []float64{}
	for _, i := range c.data {
		occurences[i]++
		if occurences[i] > c.Register.ModeRepeatCount {
			c.Register.ModeRepeatCount = occurences[i]
		}
	}
	if len(occurences) == c.length {
		c.Register.ModeRepeatCount = 0
		c.calculatedModes = true
		return c
	}
	for i, v := range occurences {
		if v == c.Register.ModeRepeatCount {
			c.Register.Modes = append(c.Register.Modes, i)
		}
	}
	c.calculatedModes = true
	return c
}

// Median calculates median of the values in Cell.
func (c *Cell) Median(sorted bool) *Cell {
	if (!sorted && c.calculatedMedian) || (sorted && c.calculatedSortedMedian) || c.err != nil {
		return c
	}
	if c.length == 1 {
		c.Register.SortedMedian = c.data[0]
		c.Register.Median = c.data[0]
		c.calculatedSortedMedian = true
		c.calculatedMedian = true
		return c
	}

	if sorted {
		if !c.even {
			c.Register.SortedMedian = c.sortedData[c.middleIndex]
		} else {
			c.Register.SortedMedian = (c.sortedData[c.middleIndex] + c.sortedData[c.middleIndex-1]) / 2
		}
		c.calculatedSortedMedian = true
	} else {
		if !c.even {
			c.Register.Median = c.data[c.middleIndex]
		} else {
			c.Register.Median = (c.data[c.middleIndex] + c.data[c.middleIndex-1]) / 2
		}
		c.calculatedMedian = true
	}
	return c
}

// StandardDeviation calculates standard deviation of the values in Cell.
func (c *Cell) StandardDeviation() *Cell {
	if c.calculatedStandardDeviation || c.err != nil {
		return c
	}

	if !c.calculatedVariance {
		c.Variance()
		if c.err != nil {
			return c
		}
	}

	c.Register.StandardDeviation = math.Sqrt(c.Register.Variance)
	c.calculatedStandardDeviation = true
	return c
}

// RunAll performs all available culculations in Cell.
func (c *Cell) RunAll() *Cell {
	if valid := c.validate(); !valid {
		return c
	}

	c.Total()
	c.StandardDeviation()
	c.Variance()
	c.Range()
	c.MaxWithIndices()
	c.MinWithIndices()
	c.Median(true)
	c.Median(false)
	c.Mean()
	c.Modes()
	return c
}

// Float64 returns float64 representation of the result
//func (c *Cell) Float64() float64 {
//	return c.Value
//}

// NewUint64 returns an instance of Cell from uint64 array.
func NewUint64(data []uint64) *Cell {
	arr := []float64{}
	for _, i := range data {
		arr = append(arr, float64(i))
	}
	return New(arr)
}

// NewUint32 returns an instance of Cell from uint32 array.
func NewUint32(data []uint32) *Cell {
	arr := []float64{}
	for _, i := range data {
		arr = append(arr, float64(i))
	}
	return New(arr)
}

// NewInt64 returns an instance of Cell from int64 array.
func NewInt64(data []int64) *Cell {
	arr := []float64{}
	for _, i := range data {
		arr = append(arr, float64(i))
	}
	return New(arr)
}

// NewInt32 returns an instance of Cell from int32 array.
func NewInt32(data []int32) *Cell {
	arr := []float64{}
	for _, i := range data {
		arr = append(arr, float64(i))
	}
	return New(arr)
}

// NewUint returns an instance of Cell from uint array.
func NewUint(data []uint) *Cell {
	arr := []float64{}
	for _, i := range data {
		arr = append(arr, float64(i))
	}
	return New(arr)
}

// NewInt returns an instance of Cell from int array.
func NewInt(data []int) *Cell {
	arr := []float64{}
	for _, i := range data {
		arr = append(arr, float64(i))
	}
	return New(arr)
}

// NewString returns an instance of Cell from string array.
func NewString(s string) *Cell {
	arr := []float64{}
	for _, i := range strings.Split(s, ",") {
		i := strings.TrimSpace(i)
		if !strings.Contains(i, ".") {
			i = i + ".0"
		}
		j, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return nil
		}
		arr = append(arr, j)
	}
	return New(arr)
}

// New returns an instance of Cell from float64 array.
// If the cell does not initialize successfully, then
// this function returns nil.
func New(data []float64) *Cell {
	if len(data) == 0 {
		return nil
	}
	c := &Cell{
		data:   data,
		length: len(data),
		even:   false,
	}

	c.sortedData = make([]float64, c.length)
	copy(c.sortedData, c.data)
	sort.Float64s(c.sortedData)
	modInt, modFrac := math.Modf(float64(c.length))
	if modFrac == 0.0 {
		c.even = true
	}
	c.middleIndex = int(modInt / 2)
	return c
}

// Failed returns true if one or more calculations failed.
func (c *Cell) Failed() bool {
	if c.err != nil {
		return true
	}
	return false
}

func addNewLines(s string, max int) string {
	if len(s) > max {
		var sb strings.Builder
		var j int
		for _, c := range s {
			if unicode.IsSpace(c) {
				if j > max {
					sb.WriteString("\n  ")
					j = 2
				}
			}
			sb.WriteRune(c)
			j++
		}
		return sb.String()
	}
	return s
}

// Print returns a string with the contents of a Cell
func (c *Cell) Print() string {
	maxWidth := 40
	var sb strings.Builder
	sb.WriteString(addNewLines(fmt.Sprintf("Data: %v", c.data), maxWidth))
	sb.WriteString(fmt.Sprintf("\nTotal: %f", c.Register.Total))
	sb.WriteString(fmt.Sprintf("\nMean: %f", c.Register.Mean))
	sb.WriteString(fmt.Sprintf("\nMedian: %f", c.Register.Median))
	sb.WriteString(fmt.Sprintf("\nSorted Median: %f", c.Register.SortedMedian))
	sb.WriteString(fmt.Sprintf("\nMax: %f", c.Register.MaxValue))
	sb.WriteString(addNewLines(fmt.Sprintf("\nMax Indices: %v", c.Register.MaxIndices), maxWidth))
	sb.WriteString(fmt.Sprintf("\nMin: %f", c.Register.MinValue))
	sb.WriteString(addNewLines(fmt.Sprintf("\nMin Indices: %v", c.Register.MinIndices), maxWidth))
	sb.WriteString(fmt.Sprintf("\nVariance: %f", c.Register.Variance))
	sb.WriteString(fmt.Sprintf("\nStandard Deviation: %f", c.Register.StandardDeviation))
	sb.WriteString(addNewLines(fmt.Sprintf("\nModes: %v", c.Register.Modes), maxWidth))
	sb.WriteString(fmt.Sprintf("\nMode Repeat Count: %d", c.Register.ModeRepeatCount))

	return sb.String()
}
