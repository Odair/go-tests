package main

import "fmt"

const prefixoOlaPortugues = "Olá "
const prefixoOlaFrances = "Bonjour "
const prefixoOlaEspanhol = "Hola "

func prefixodeSaudacao(idioma string) (prefixo string) {
    switch idioma {
    case "frances":
        prefixo = prefixoOlaFrances
    case "espanhol":
        prefixo = prefixoOlaEspanhol
    default:
        prefixo = prefixoOlaPortugues
    }
    return
}

func Hello(name string, idioma string) string {
	if name == "" {
        name = "mundo"
    }
	return prefixodeSaudacao(idioma) + name
}

func main() {
	fmt.Println(Hello("teste", ""))
}