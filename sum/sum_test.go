package ints

import "testing"

func TestSum(t *testing.T) {
	soma := Adiciona(2, 2)
    esperado := 4

    if soma != esperado {
        t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
    }
}