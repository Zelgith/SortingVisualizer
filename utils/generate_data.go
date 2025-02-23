package utils

import (
	"math/rand"
	"time"
)

// GenerateIntData generates a slice of random integers of the specified length.
func GenerateIntData(length, maxValue int) []any {
	rand.Seed(time.Now().UnixNano())
	data := make([]any, length)
	for i := 0; i < length; i++ {
		data[i] = rand.Intn(maxValue) + 1
	}
	return data
}

// GenerateFloatData generates a slice of random floats of the specified length.
func GenerateFloatData(length, maxValue int) []any {
	rand.Seed(time.Now().UnixNano())
	data := make([]any, length)
	for i := 0; i < length; i++ {
		data[i] = rand.Float64()*float64(maxValue) + 1.0
	}
	return data
}
