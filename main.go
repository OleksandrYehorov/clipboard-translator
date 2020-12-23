package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	mTranslateDomain := "www.m-translate.com"
	mTranslateURL := fmt.Sprintf("https://%s", mTranslateDomain)
	// fromLang := "en"

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), chromedp.Headless)
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 1390*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var translatedText string
	var evalOutput string

	err := chromedp.Run(ctx,
		chromedp.Navigate(mTranslateURL),
		chromedp.WaitVisible("#text"),
		chromedp.Evaluate(`(() => {
			document.querySelector('#translate_from').value = 'ru';
			document.querySelector('#translate_to').value = 'en';
			document.querySelector('#text').value = 'Как дела?';
			document.querySelector('#go_btn').click();
		})() || ''`, &evalOutput),
		chromedp.WaitNotVisible("#out_cont .loading"),
		chromedp.Evaluate("document.querySelector('#text_out').value", &translatedText),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Translation: %s", strings.TrimSpace(translatedText))

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
