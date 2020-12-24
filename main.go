package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"

	"github.com/OleksandrYehorov/clipboard-translator/mtranslator"
	"github.com/OleksandrYehorov/clipboard-translator/translator"
)

func main() {
	log.SetFlags(0)

	var t translator.Translator
	var cancel context.CancelFunc

	t, cancel = mtranslator.New()
	defer cancel()

	textsToTranslate := []string{
		"травень",
	}

	fromLang, toLang := "uk", "en"

	for _, text := range textsToTranslate {
		translatedtext, error := t.Translate(fromLang, toLang, text)
		if error != nil {
			log.Fatal(error)
		}
		log.Printf("%v -> %v", text, strings.TrimSpace(translatedtext))
	}

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
