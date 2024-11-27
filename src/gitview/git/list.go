package git

import (
	"fmt"
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
	fmt.Println("Changed files:")
	for file, fileStatus := range status {
		fmt.Printf("- %s: %s\n", file, string(fileStatus.Worktree)) // For debugging, to be removed
		fileList = append(fileList, file)
		filesStatusList = append(filesStatusList, string(fileStatus.Worktree))
	}
	return fileList, filesStatusList
}