package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func extractData(accessToken string) string {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlContent string

	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://progress.alem.school/`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			err := chromedp.Evaluate(`localStorage.setItem('access_token', '`+accessToken+`')`, nil).Do(ctx)
			if err != nil {
				return err
			}
			fmt.Println("Access token set in local storage")
			return nil
		}),
		chromedp.Reload(),
		chromedp.Sleep(3*time.Second),
		chromedp.WaitVisible(`#root`, chromedp.ByID),
		chromedp.OuterHTML(`#root`, &htmlContent),
	)

	if err != nil {
		log.Fatal(err)
	}

	return htmlContent
}

func extractDataHelper(content, key string) string {
	startIdx := strings.Index(content, key)
	if startIdx == -1 {
		return ""
	}

	startIdx += len(key)

	endIdx := strings.Index(content[startIdx:], "</p>")
	if endIdx == -1 {
		return ""
	}

	value := strings.TrimSpace(content[startIdx : startIdx+endIdx])
	return value
}

func printExtracted(htmlContent string) {
	login := extractDataHelper(htmlContent, "Login:")
	hours := extractDataHelper(htmlContent, "Hours:")
	Life := extractDataHelper(htmlContent, "Life:")

	fmt.Println("\n===================================")
	fmt.Println("Login:", login)
	fmt.Println("Hours:", hours)
	fmt.Println("Life:", Life)
	fmt.Println("===================================")

	timeChecker(hours)
}
