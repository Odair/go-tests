package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, expected string) {
        t.Helper()
        if resultado != expected {
            t.Errorf("resultado '%s', esperado '%s'", resultado, expected)
        }
    }

	t.Run("say hello to everyone", func (t *testing.T) {
		resultado := Hello("Odair", "")
		expected := "Olá Odair"

		verificaMensagemCorreta(t, resultado, expected)
	})

	t.Run("say 'Olá mundo' when string is empty", func (t *testing.T) {
		resultado := Hello("", "")
		expected := "Olá mundo"

		verificaMensagemCorreta(t, resultado, expected)
	})

	t.Run("em espanhol", func(t *testing.T) {
        resultado := Hello("Elodie", "espanhol")
        esperado := "Hola Elodie"
        verificaMensagemCorreta(t, resultado, esperado)
    })
}