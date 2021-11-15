package functions

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var DirPath = "empty"
var BasePath = "empty"
var mark = 0

// A funtion to check if the provided directory is a valid must-gather diretory
func DirCheck() {
	// Check if there is a provided argument
	len := len(os.Args)
	if len < 2 {
		// Check if this argument is a directory
		fmt.Print("Please provide the must-gather directory" + "\n")
	} else {
		DirPath = os.Args[1]
		// Try to acces this Directory
		file, err := os.Open(DirPath)
		if err != nil {
			fmt.Print("I couldn't access this directory! .. Please check if you have the correct permissions" + "\n")
		} else {
			fileInfo, err := file.Stat()
			if err != nil {
				fmt.Print("Something went wrong!" + "\n")
			} else {
				if fileInfo.IsDir() {
					// Check if this is a valid must-gather direcotry
					err := filepath.Walk(DirPath,
						func(path string, info os.FileInfo, err error) error {
							if err != nil {
								return err
							}
							if info.Name() == "cluster-scoped-resources" {
								mark = mark + 1
							} else if info.Name() == "namespaces" {
								mark = mark + 1
							}
							return nil
						})
					if mark < 2 {
						fmt.Print("This is not a valid must-gather directory" + "\n")
					} else {
						// If we are sure this is a must-gather director, go get the base path "the directory path containing the MG subdirectories"
						DirPathTreeBuild()
					}
					if err != nil {
						log.Println(err)
					}
				} else {
					fmt.Print("This is a not a Directory!" + "\n")
				}
			}
		}

	}
}

func DirPathTreeBuild() {
	// Get the path base. For example: ../a/b/c/must-gather.local.nnn/quay-io-openshift-release-xxx/
	err := filepath.Walk(DirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "cluster-scoped-resources" {
				mypath := strings.Split(path, "cluster-scoped-resources")
				BasePath = mypath[0]
				// Start the MGR interface
				StartUI()
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func StartUI() {
	// Call function CreateUI from functions/initialUI.go package to create the initial UI
	// Sart the UI with the Root=MainGrid
	CreateUI()
	App.SetInputCapture(KeyboardKeys).
		SetRoot(MainGrid, true).
		EnableMouse(true).
		Run()

}
