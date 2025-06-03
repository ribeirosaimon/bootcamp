package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	for _, v := range []struct {
		minutes uint
		cat     Categoria
		want    uint
	}{
		{
			minutes: 120,
			cat:     CatA,
			want:    9000,
		},
		{
			minutes: 60,
			cat:     CatB,
			want:    1725,
		},
		{
			minutes: 180,
			cat:     CatC,
			want:    3000,
		},
	} {
		t.Run(fmt.Sprintf("Deve retornar %d\n", v.want), func(t *testing.T) {
			percentage := calcSalario(v.minutes, v.cat)
			if percentage != v.want {
				t.Errorf("Erro")
			}
		})
	}
}
