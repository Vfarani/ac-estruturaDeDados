package main

import (
	"fmt"
	"strings"
)

func main() {
	frasesNatal := map[string]string{
		"alemanha":       "Frohliche Weihnachten!",
		"austria":        "Frohe Weihnacht!",
		"suecia":         "God Jul!",
		"antardida":      "Merry Christmas!",
		"grecia":         "Kala Christougena!",
		"estados-unidos": "Merry Christmas!",
		"inglaterra":     "Merry Christmas!",
		"australia":      "Merry Christmas!",
		"brasil":         "Feliz Natal!",
		"portugal":       "Feliz Natal!",
		"espanha":        "Feliz Navidad!",
		"turquia":        "Mutlu Noeller",
		"argentina":      "Feliz Navidad!",
		"chile":          "Feliz Navidad!",
		"japao":          "Merii Kurisumasu!",
		"mexico":         "Feliz Navidad!",
		"canada":         "Merry Christmas!",
		"irlanda":        "Nollaig Shona Dhuit!",
		"belgica":        "Zalig Kerstfeest!",
		"italia":         "Buon Natale!",
		"libia":          "Buon Natale!",
		"siria":          "Milad Mubarak!",
		"marrocos":       "Milad Mubarak!",
		"coreia":         "Chuk Sung Tan!",

	}

	paises := []string{"uri-online-judge", "alemanha", "brasil", "austria", "china"}

	for i, pais := range paises {
		if i != 0 {
			fmt.Println()
		}
		frase, ok := frasesNatal[strings.ToLower(pais)]
		if ok {
			fmt.Println(frase)
		} else {
			fmt.Println("Não Encontrado!")
		}
	}
}
