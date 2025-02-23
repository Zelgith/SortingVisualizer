package utils

import (
	"context"
)

type Sorter interface {
	Sort([]any, func([]any, []int))
}

type SortTemplate struct {
	Slice     *[]any
	Algorithm Sorter
	DrawFunc  func([]any, []int)
}

func (st *SortTemplate) SortWithContext(ctx context.Context) {
	// Sort the slice
	st.Algorithm.Sort(*st.Slice, func(data []any, step []int) {
		select {
		case <-ctx.Done():
			return
		default:
			st.DrawFunc(*st.Slice, step)
		}
	})

	// Highlight the sorted slice
	step := []int{}
	for i := 0; i < len(*st.Slice); i++ {
		select {
		case <-ctx.Done():
			return
		default:
			step = append(step, i)
			st.DrawFunc(*st.Slice, step)
		}
	}
}

// Compare function to handle comparisons between different data types
func Compare(a, b any) int {
	switch a := a.(type) {
	case int:
		return a - b.(int)
	case float64:
		b := b.(float64)
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	default:
		return 0
	}
}
