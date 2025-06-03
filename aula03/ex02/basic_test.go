package main

import (
	"testing"
)

func TestName(t *testing.T) {
	t.Run("Deve retornar 3.0\n", func(t *testing.T) {
		media := calcMedia(1, 2, 3, 4, 5)
		if media != 3 {
			t.Errorf("Erro")
		}
	})
}
