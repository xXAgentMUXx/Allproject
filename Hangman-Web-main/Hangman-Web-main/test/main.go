package main

import (
	"fmt"
	"unicode"
)

func cleanText(text string) string {
	result := ""
	for _, r := range text {
		// On conserve les lettres, chiffres, espaces et signes de ponctuation,
		// en excluant seulement les caractères non imprimables.
		if unicode.IsPrint(r) {
			result += string(r)
		}
	}
	return result
}

func main() {
	texte := "école"
	texteNettoye := cleanText(texte)
	fmt.Println("Texte nettoyé :", texteNettoye)
}