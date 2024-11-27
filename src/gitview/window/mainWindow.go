package window

import (
	"fmt"
	"gitview/src/gitview/git"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var amend bool

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

// TODO: Actually add the list of changes to stage
func gitTopHalf() *fyne.Container {
	changesLabel := widget.NewLabel("Changes:")
	amendCheckBox := widget.NewCheck("Amend", func(b bool) {
		if(b){
			amend = true
		}else{
			amend = false
		}
	})
	content := container.NewVBox(changesLabel, listChanges(), amendCheckBox)
	return content
}

func gitBottomHalf() *fyne.Container {
	commitMessageEntry := widget.NewEntry()
	commitButton := widget.NewButton("Commit", func() {git.Commit(commitMessageEntry.Text, amend)})
	commitPushButton := widget.NewButton("Commit and Push", func() {git.CommitAndPush(commitMessageEntry.Text, amend)})
	buttonHBox := container.NewHBox(commitButton, commitPushButton)
	content := container.NewVBox(commitMessageEntry, buttonHBox)
	return content
}

func gitContent() *container.Split {
	content := container.NewVSplit(gitTopHalf(), gitBottomHalf())
	return content
}

func LaunchMainWindow(a fyne.App) {
	windowContent := container.NewVBox(topBar(), gitContent())
	mainWindow := a.NewWindow("GitView")
	mainWindow.SetContent(windowContent)
	mainWindow.Resize(fyne.NewSize(350,800))
	mainWindow.Show()
}