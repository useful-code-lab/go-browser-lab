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
		Description: "Простой браузер на Go и Wails",

		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},

		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Главное окно браузера.
	mainWindow := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:             "main",
		Title:            "Go Browser",
		Width:            1200,
		Height:           800,
		MinWidth:         800,
		MinHeight:        600,
		BackgroundColour: application.NewRGB(30, 30, 30),
		URL:              "/",
	})

	// Окно для отображения сайтов.
	pageWindow := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:             "page",
		Title:            "Go Browser",
		Width:            1200,
		Height:           700,
		MinWidth:         800,
		MinHeight:        500,
		Hidden:           true,
	})

	_ = mainWindow
	_ = pageWindow

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}