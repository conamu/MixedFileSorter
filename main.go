package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	currentPath string
	err         error
)

func main() {
	if len(os.Args) < 2 {
		// Get Current path to work in
		currentPath, err = os.Getwd()
		handleError(err)
	} else {
		currentPath = os.Args[1]
	}
	log.Println("Working in current directory: " + currentPath)

	folders, err := os.ReadDir(currentPath)
	handleError(err)

	log.Println("Processing these directories: ")
	for _, dir := range folders {
		if dir.IsDir() {
			fmt.Println(dir.Name())
		}
	}

	// Range over every folder
	for _, dir := range folders {
		// If no directory, skip this entry
		if !dir.IsDir() {
			continue
		}

		pathOfDirectory := currentPath + dir.Name()
		log.Println("Processing directory " + dir.Name())

		innerFiles, err := os.ReadDir(pathOfDirectory)
		handleError(err)

		// Get info of the File to Copy
		for _, innerFile := range innerFiles {
			if innerFile.IsDir() {
				continue
			}
			innerFolders := strings.Split(dir.Name(), ", ")
			log.Println("Splitting folder " + dir.Name() + " into " + strings.Join(innerFolders, " "))
			log.Println("Processing file " + innerFile.Name())

			// Copy the file into the split folders
			for _, directoryToPutStuffIn := range innerFolders {
				pathForFile := currentPath + "orderedFiles/" + directoryToPutStuffIn + "/"
				filePath := pathForFile + directoryToPutStuffIn + innerFile.Name()

				if err := os.MkdirAll(pathForFile, 0755); err != nil {
					log.Fatal(err.Error())
				}
				pathOfFileToCopy := pathOfDirectory + "/" + innerFile.Name()
				handleError(copyFileToFolder(pathOfFileToCopy, filePath))

			}

		}

	}

}

func copyFileToFolder(source, destination string) error {

	input, err := ioutil.ReadFile(source)
	if err != nil {
		log.Println("Error reading", source)
		return err
	}

	err = ioutil.WriteFile(destination, input, 0644)
	if err != nil {
		log.Println("Error creating", destination)
		return err
	}

	return nil
}

func handleError(err error) {
	if err != nil {
		log.Fatalln("Program encountered an error " + err.Error())
	}
}
