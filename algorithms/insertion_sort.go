package algorithms

import "github.com/Zelgith/SortingVisualizer/utils"

type InsertionSort struct{}

func (InsertionSort) Sort(slice []any, drawBarChart func([]any, []int)) {
	for i := 1; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && utils.Compare(slice[j], key) > 0 {
			slice[j+1] = slice[j]
			j--
			drawBarChart(slice, []int{j + 1, j + 2})
		}
		slice[j+1] = key
		drawBarChart(slice, []int{j + 1})
	}
}
