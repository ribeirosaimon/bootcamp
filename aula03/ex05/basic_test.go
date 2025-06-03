package main

import (
	"testing"
)

func TestTarantula(t *testing.T) {
	for _, v := range []struct {
		name     string
		value    int
		expected float64
		auxFunc  func(int) float64
	}{
		{
			name:     "Tarantula",
			expected: 15,
			value:    100,
			auxFunc:  getTarantula,
		},
		{
			name:     "Dog",
			expected: 100,
			value:    10,
			auxFunc:  getDog,
		},
		{
			name:     "Cat",
			expected: 50,
			value:    10,
			auxFunc:  getCat,
		},
		{
			name:     "Hamster",
			expected: 25,
			value:    100,
			auxFunc:  getHamster,
		},
	} {
		t.Run(v.name, func(t *testing.T) {
			res := v.auxFunc(v.value)
			if res != v.expected {
				t.Errorf("Expected %f, got %f", v.expected, res)
			}
		})
	}
}
