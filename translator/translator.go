package translator

// Translator interface
type Translator interface {
	Translate(fromLang string, toLang string, textToTranslate string) (translatedText string, error error)
}
