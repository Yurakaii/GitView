package git

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
)

func ListChanges() ([]string, []string){
	// Open the repository
	repo, err := git.PlainOpen(".") // Pull directory from config
	if err != nil {
		log.Fatalf("Failed to open repository: %v", err)
	}

	// Get the worktree
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Failed to get worktree: %v", err)
	}

	// Get the status of the working directory
	status, err := worktree.Status()
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