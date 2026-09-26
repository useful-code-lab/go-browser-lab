package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	webview "github.com/moxcomic/go-webview"
	_ "github.com/moxcomic/go-webview/embedded"
)

const browserHTML = `
<!DOCTYPE html>
<html lang="ru">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">

	<title>Go Browser</title>

	<style>
		* {
			box-sizing: border-box;
		}

		html,
		body {
			margin: 0;
			padding: 0;
			width: 100%;
			height: 100%;
			overflow: hidden;
			font-family: Arial, sans-serif;
			background: #1e1e1e;
		}

		body {
			display: flex;
			flex-direction: column;
		}

		.toolbar {
			display: flex;
			gap: 8px;
			padding: 8px;
			background: #252526;
			border-bottom: 1px solid #3e3e42;
		}

		#url {
			flex: 1;
			height: 36px;
			padding: 0 12px;

			border: 1px solid #555;
			border-radius: 6px;

			background: #1e1e1e;
			color: #fff;

			font-size: 14px;
			outline: none;
		}

		#url:focus {
			border-color: #0078d4;
		}

		button {
			height: 36px;
			padding: 0 18px;

			border: none;
			border-radius: 6px;

			background: #0078d4;
			color: white;

			cursor: pointer;
			font-size: 14px;
		}

		button:hover {
			background: #106ebe;
		}

		#browser {
			flex: 1;

			display: flex;
			align-items: center;
			justify-content: center;

			color: #aaa;
			font-size: 16px;
		}
	</style>
</head>

<body>

	<div class="toolbar">

		<input
			id="url"
			type="text"
			value="https://google.com"
			placeholder="Введите адрес сайта"
		>

		<button onclick="navigate()">
			Перейти
		</button>

	</div>

	<div id="browser">
		Введите адрес сайта
	</div>

	<script>

		function normalizeURL(value) {

			value = value.trim();

			if (!value) {
				return "";
			}

			if (!/^https?:\/\//i.test(value)) {
				value = "https://" + value;
			}

			return value;
		}

		function navigate() {

			const input = document.getElementById("url");

			const url = normalizeURL(input.value);

			if (!url) {
				return;
			}

			window.location.href = url;
		}

		document
			.getElementById("url")
			.addEventListener("keydown", function(event) {

				if (event.key === "Enter") {
					navigate();
				}

			});

	</script>

</body>
</html>
`

func normalizeURL(rawURL string) (string, error) {
	rawURL = strings.TrimSpace(rawURL)

	if rawURL == "" {
		return "", fmt.Errorf("URL пустой")
	}

	if !strings.HasPrefix(rawURL, "http://") &&
		!strings.HasPrefix(rawURL, "https://") {

		rawURL = "https://" + rawURL
	}

	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return "", fmt.Errorf("некорректный URL: %w", err)
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", fmt.Errorf("поддерживаются только HTTP и HTTPS")
	}

	return parsed.String(), nil
}

func main() {
	wv, err := webview.New(true)
	if err != nil {
		log.Fatalf("не удалось создать WebView: %v", err)
	}

	defer wv.Destroy()

	wv.SetTitle("Go Browser")
	wv.SetSize(1200, 800, webview.HintNone)

	wv.SetHtml(browserHTML)

	wv.Run()
}