package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Mau005/MyExpenses/db"
)

var W float32 = 800
var H float32 = 600

func main() {
	db.ConnectionSqlite()
	myApp := app.New()
	myWindow := myApp.NewWindow("Wallet")

	// Crear elementos de la interfaz
	balanceLabel := widget.NewLabel("Balance: $100")
	sendButton := widget.NewButton("Enviar", func() {
		// Lógica para enviar dinero
	})
	receiveButton := widget.NewButton("Recibir", func() {
		// Lógica para recibir dinero
	})

	// Crear el diseño de la interfaz
	content := container.NewVBox(
		balanceLabel,
		sendButton,
		receiveButton,
	)

	// Configurar la ventana
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}

func Precion() {
	fmt.Println("Estoy siendo precionado")
}
