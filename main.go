package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 1. Создаем приложение и главное окно
	myApp := app.New()
	myWindow := myApp.NewWindow("Мой Go Браузер")

	// 2. Создаем элементы интерфейса
	addressEntry := widget.NewEntry()
	addressEntry.SetPlaceHolder("Введите URL (например, https://google.com)")

	goButton := widget.NewButton("Перейти", func() {
		url := addressEntry.Text
		fmt.Printf("Переход по адресу: %s\n", url)
		// Здесь позже будет логика загрузки страницы
	})

	// 3. Верстка: размещаем адресную строку и кнопку в один ряд (HBox)
	topBar := container.NewBorder(nil, nil, nil, goButton, addressEntry)

	// Временный контент для центральной части
	mainContent := widget.NewLabel("Здесь будет отображаться веб-страница")

	// Комнуем всё вместе: панель управления сверху, контент по центру
	content := container.NewBorder(topBar, nil, nil, nil, mainContent)

	// 4. Запуск окна
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(1024, 768))
	myWindow.ShowAndRun()
}
