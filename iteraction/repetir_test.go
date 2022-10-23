package iteracao

import "testing"

func TestRepetir(t *testing.T) {
    repeticoes := Repetir("a", 5)
    esperado := "aaaaa"

    if repeticoes != esperado {
        t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
    }
}
