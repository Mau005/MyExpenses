package handler

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"a", "string", "list"}

func LobbyHandler(window fyne.Window) {
	myName := widget.NewLabel("Mauricio Pino")
	myWorks := widget.NewLabel("Actriz Porno")
	mySalary := widget.NewLabel("11111111$")
	contentLayout := container.New(layout.NewVBoxLayout(), myName, myWorks, mySalary)
	contentLayout2 := container.New(layout.NewCenterLayout(), contentLayout)
	card := widget.NewCard("Mis Estadisticas", "Etc", contentLayout2)

	window.SetContent(card)

}
