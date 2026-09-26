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

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:              "browser",
		Title:             "Go Browser",
		Width:             1200,
		Height:            800,
		MinWidth:          800,
		MinHeight:         600,
		BackgroundColour:  application.NewRGB(30, 30, 30),
		URL:               "/",
		DevToolsEnabled:   true,
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}