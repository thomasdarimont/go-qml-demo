package main

import (
	"os"

	"github.com/go-qamel/qamel"
)

func init() {
	// Register the BackEnd as QML component
	RegisterQmlBackEnd("BackEnd", 1, 0, "BackEnd")
}

// #cgo LDFLAGS: -L /usr/include/qt5/QtWidgets
func main() {
	// Create application
	app := qamel.NewApplication(len(os.Args), os.Args)
	app.SetApplicationDisplayName("Qamel Example")

	// Create a QML viewer
	view := qamel.NewViewerWithSource("qrc:/res/ui.qml")
	view.SetResizeMode(qamel.SizeRootObjectToView)
	view.SetHeight(300)
	view.SetWidth(400)
	view.Show()

	// Exec app
	app.Exec()
}
