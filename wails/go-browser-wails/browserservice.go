package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type BrowserService struct {
	window *application.WebviewWindow
}

func NewBrowserService(window *application.WebviewWindow) *BrowserService {
	return &BrowserService{
		window: window,
	}
}

// Navigate подготавливает URL.
// На первом этапе навигация выполняется через основное WebView.
func (b *BrowserService) Navigate(rawURL string) error {
	rawURL = strings.TrimSpace(rawURL)

	if rawURL == "" {
		return nil
	}

	// Если пользователь ввёл поисковый запрос,
	// пока отправляем его в Google.
	if !strings.Contains(rawURL, "://") &&
		!strings.Contains(rawURL, ".") {

		rawURL = "https://www.google.com/search?q=" +
			url.QueryEscape(rawURL)
	}

	if !strings.Contains(rawURL, "://") {
		rawURL = "https://" + rawURL
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("некорректный адрес: %w", err)
	}

	if parsedURL.Scheme != "http" &&
		parsedURL.Scheme != "https" {
		return fmt.Errorf("поддерживаются только HTTP и HTTPS")
	}

	b.window.SetURL(parsedURL.String())

	return nil
}

func (b *BrowserService) Reload() {
	b.window.Reload()
}

func (b *BrowserService) Back() {
	// Реализацию истории добавим следующим этапом.
}

func (b *BrowserService) Forward() {
	// Реализацию истории добавим следующим этапом.
}

func (b *BrowserService) Home() {
	b.Navigate("https://www.google.com")
}