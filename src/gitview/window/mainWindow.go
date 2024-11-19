package window

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// TODO: Fix this
func createSmallIcon(resource fyne.Resource) *canvas.Image {
	icon := canvas.NewImageFromResource(resource)
	icon.Resize(fyne.NewSize(16,16))
	return icon
}

func topBar() *fyne.Container {
	// TODO: Move the internal functions and add git integration.
	// TODO: Make a function that makes the button variables to avoid recursion.
	reloadButton := widget.NewButtonWithIcon("", createSmallIcon(theme.ViewRefreshIcon()).Resource, func() {
		fmt.Println("Reload Button Pressed") 
	})
	undoButton := widget.NewButtonWithIcon("", createSmallIcon(theme.ContentUndoIcon()).Resource, func() {
		fmt.Println("Undo Button Pressed")
	})
	downloadButton := widget.NewButtonWithIcon("", createSmallIcon(theme.DownloadIcon()).Resource, func() {
		fmt.Println("Download Button Pressed")
	})
	visibilityButton := widget.NewButtonWithIcon("", createSmallIcon(theme.VisibilityIcon()).Resource, func() {
		fmt.Println("Visibilty Button Pressed")
	})
	expandButton := widget.NewButtonWithIcon("", createSmallIcon(theme.ZoomInIcon()).Resource, func() {
		fmt.Println("Expand Button Pressed")
	})
	collapseButton := widget.NewButtonWithIcon("", createSmallIcon(theme.ZoomOutIcon()).Resource, func() {
		fmt.Println("Collapse Button Pressed")
	})
	toolbar := container.NewHBox(
		reloadButton,
		undoButton,
		downloadButton,
		visibilityButton,
		expandButton,
		collapseButton,
	)
	return toolbar
}

func LaunchWindow(a fyne.App) {
	content := container.NewVBox(topBar())
	mainWindow := a.NewWindow("GitView")
	mainWindow.SetContent(content)
	mainWindow.Resize(fyne.NewSize(350,800))
	mainWindow.Show()
}