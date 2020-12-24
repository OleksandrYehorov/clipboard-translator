package mtranslator

import (
	"context"
	"log"
	"time"

	cdp "github.com/chromedp/chromedp"
)

// MTranslator struct
type MTranslator struct {
	ctx context.Context
}

// New creates an instance of MTranslator
func New() (t MTranslator, cancel context.CancelFunc) {
	t.ctx, cancel = cdp.NewExecAllocator(context.Background(), cdp.Headless)
	t.ctx, cancel = cdp.NewContext(
		t.ctx,
		cdp.WithLogf(log.Printf),
	)

	return
}

// Translate translates text using m-translate.com
func (t MTranslator) Translate(fromLang string, toLang string, textToTranslate string) (translatedText string, error error) {
	error = cdp.Run(t.ctx, cdp.Tasks{
		cdp.Navigate("https://www.m-translate.com"),
		cdp.WaitVisible("#text", cdp.ByID),
		cdp.Sleep(100 * time.Millisecond), // need this to work in headless mode
		cdp.SetValue("#translate_from", fromLang, cdp.ByID),
		cdp.SetValue("#translate_to", toLang, cdp.ByID),
		cdp.SetValue("#text", textToTranslate, cdp.ByID),
		cdp.Click("#go_btn", cdp.ByID),
		cdp.WaitVisible("#out_cont .loading", cdp.ByQuery),
		cdp.WaitNotVisible("#out_cont .loading", cdp.ByQuery),
		cdp.Value("#text_out", &translatedText, cdp.ByID),
	})

	return
}
