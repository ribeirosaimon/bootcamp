package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	for _, v := range []struct {
		salario uint
		want    uint
	}{
		{
			salario: 40,
			want:    0,
		},
		{
			salario: 60,
			want:    17,
		},
		{
			salario: 200,
			want:    27,
		},
	} {
		t.Run(fmt.Sprintf("Deve retornar %d\n", v.want), func(t *testing.T) {
			percentage := getPercentage(v.salario)
			if percentage != v.want {
				t.Errorf("Erro")
			}
		})
	}
}
