package git

import (
	"log"
)

func ListChanges() ([]string, []string){
	// Get the status of the working directory
	status, err := Worktree.Status()
	if err != nil {
		log.Fatalf("Failed to get status: %v", err)
	}
	var fileList, filesStatusList []string
	// Iterate through the status to find changed files
	for file, fileStatus := range status {
		fileList = append(fileList, file)
		filesStatusList = append(filesStatusList, string(fileStatus.Worktree))
	}
	return fileList, filesStatusList
}