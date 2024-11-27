package window

import (
	"gitview/src/gitview/git"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func listChanges() *fyne.Container{
	fileList, filesStatusList := git.ListChanges()
	length := len(fileList)
	listVBox := container.NewVBox()
	for i := 0; i < length; i++ {
		file, fileStatus := fileList[i], filesStatusList[i]
		fileLabel, fileStatusLabel := widget.NewLabel(file), widget.NewLabel(fileStatus)
		stageButton := widget.NewButton("+", func() {git.Stage(file)})
		fileHBox := container.NewHBox(fileLabel, fileStatusLabel, stageButton)
		listVBox.Add(fileHBox)
	}
	return listVBox
}