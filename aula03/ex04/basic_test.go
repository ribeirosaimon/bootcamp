package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	for _, v := range []struct {
		want    float64
		auxFunc func(...float64) float64
	}{
		{
			want:    1,
			auxFunc: minValue,
		},
		{
			want:    5,
			auxFunc: maxValue,
		},
		{
			want:    3,
			auxFunc: averageValue,
		},
	} {
		t.Run(fmt.Sprintf("Deve retornar %v\n", v.want), func(t *testing.T) {
			result := v.auxFunc(1, 2, 3, 4, 5)
			if result != v.want {
				t.Errorf("Erro")
			}
		})
	}
}
