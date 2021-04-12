package generator

import "fmt"

type SliceIntGenerator struct {
	data    []int
	current int
}

func NewSliceIntGenerator(slices []int) *SliceIntGenerator {
	return &SliceIntGenerator{
		data:    slices,
		current: 0,
	}
}

func (g *SliceIntGenerator) Next() (int, error) {
	if g.current+1 >= len(g.data) {
		return 0, fmt.Errorf("No more elements")
	}

	g.current++
	return g.data[g.current-1], nil
}
