package mtranslator

import (
	"context"
	"testing"

	"github.com/OleksandrYehorov/clipboard-translator/translator"
)

func TestTranslate(test *testing.T) {
	var t translator.Translator
	var cancel context.CancelFunc

	t, cancel = New()
	defer cancel()

	cases := []struct {
		fromLang        string
		toLang          string
		textToTranslate string
		expected        string
	}{
		{"auto", "en", "травень", "May"},
		{"en", "uk", "a mouse", "миша"},
		{"auto", "uk", "a worm", "черв'як"},
		{"ru", "en", "книга", "book"},
	}

	for _, testCase := range cases {
		translatedtext, error := t.Translate(testCase.fromLang, testCase.toLang, testCase.textToTranslate)
		if error != nil {
			test.Error(error)
		}
		if translatedtext != testCase.expected {
			test.Errorf("Expected: %v Actual: %v", testCase.expected, translatedtext)
		}
	}
}
