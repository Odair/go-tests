package iteracao


func Repetir(caractere string, quantidadeRepeticoes int) string {
    var repeticoes string
    for i := 0; i < quantidadeRepeticoes; i++ {
        repeticoes = repeticoes + caractere
    }
    return repeticoes
}