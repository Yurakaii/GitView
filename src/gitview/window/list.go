package window

import (
	"gitview/src/gitview/git"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func listChanges() *fyne.Container{
	fileList, filesStatusList := git.ListChanges()
	length := len(fileList)
	listVBox := container.NewVBox()
	for i := 0; i < length; i++ {
		file, fileStatus := fileList[i], filesStatusList[i]
		fileLabel, fileStatusLabel := widget.NewLabel(file), widget.NewLabel(fileStatus)
		fileHBox := container.NewHBox(fileLabel, layout.NewSpacer(), fileStatusLabel)
		stageButton := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
			git.Stage(file)
			loadContent()
		})
		unStageButton := widget.NewButtonWithIcon("", theme.ContentRemoveIcon(), func() {
			git.UnStage(file)
			loadContent()
		})
		if (fileStatus == " ") {
			fileHBox.Add(unStageButton)
		} else {
			fileHBox.Add(stageButton)
		}
		listVBox.Add(fileHBox)
	}
	return listVBox
}