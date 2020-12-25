package main

import (
	"context"
	"log"

	"github.com/OleksandrYehorov/clipboard-translator/mtranslator"
	"github.com/OleksandrYehorov/clipboard-translator/translator"
	"github.com/atotto/clipboard"

	hook "github.com/robotn/gohook"
)

func main() {
	hook.Register(hook.KeyDown, []string{"ctrl", "c"}, func(e hook.Event) {
		log.Print(e.Clicks)
		var t translator.Translator
		var cancel context.CancelFunc

		textToTranslate, _ := clipboard.ReadAll()

		t, cancel = mtranslator.New()
		defer cancel()
		fromLang, toLang := "auto", "ru"
		translatedtext, _ := t.Translate(fromLang, toLang, textToTranslate)

		clipboard.WriteAll(translatedtext)
	})

	s := hook.Start()
	<-hook.Process(s)
}
