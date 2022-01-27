package functions

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GetMGFiles() {
	// Select directories under the probided MG folder and try to extract the type of each one [default, Cluseter logging, Service Mesh,...]
	Args := len(os.Args)
	if Args < 2 {
		// Check if user procided an argument [should be the MG directory]
		fmt.Print("Please provide the must-gather directory while runnnig this tool" + "\n")
	} else {
		ProvidedDirPath = os.Args[1]
		// Try to open the provided directory
		_, err := os.Open(ProvidedDirPath)
		if err != nil {
			fmt.Print("I couldn't access this directory! .. Please check if you have the correct permissions to do so" + "\n")
		} else {
			MGFiles, err := ioutil.ReadDir(ProvidedDirPath)
			if err != nil {
				fmt.Print("I couldn't read the contents of the this directory" + "\n")
			} else if len(MGFiles) > 0 {
				// Get must-gather image direcoty name and add it to the MGDropDown
				for i := range MGFiles {
					if MGFiles[i].IsDir() {
						MGDropDown.AddOption(MGFiles[i].Name(), MGDropDownOnSelect)
					}
				}
			}
			StartUI()
		}
	}
}

func StartUI() {
	// Will call function [CreateUI] from [functions/initialUI.go] package to create the initial UI
	// Start the UI with the [Root=MainGrid]
	CreateUI()
	App.SetInputCapture(KeyboardKeys).
		SetRoot(MainGrid, true).
		EnableMouse(true).
		Run()

}
