package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        "Go Browser",
		Description: "Полноценный браузер на Go и Wails",

		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},

		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Главное окно с интерфейсом браузера.
	mainWindow := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:             "browser-ui",
		Title:            "Go Browser",
		Width:            1200,
		Height:           800,
		MinWidth:         900,
		MinHeight:        600,
		BackgroundColour: application.NewRGB(20, 20, 22),
		URL:              "/",
		DevToolsEnabled:  true,
	})

	app.RegisterService(
		application.NewService(
			NewBrowserService(mainWindow),
		),
	)

	mainWindow.Show()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}