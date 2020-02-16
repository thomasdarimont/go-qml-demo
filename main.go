package main

import (
	"os"

	"github.com/go-qamel/qamel"
)

func init() {
	// Register the BackEnd as QML component
	// Note: RegisterQmlBackEnd is generated by qamel build ...
	RegisterQmlBackEnd("BackEnd", 1, 0, "BackEnd")
}

func main() {

	app := qamel.NewApplication(len(os.Args), os.Args)
	app.SetApplicationDisplayName("Qamel Example")

	view := qamel.NewViewerWithSource("qrc:/res/ui.qml")
	view.SetResizeMode(qamel.SizeRootObjectToView)
	view.SetHeight(300)
	view.SetWidth(400)
	view.Show()

	app.Exec()
}
