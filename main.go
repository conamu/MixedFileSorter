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

	// Read folderstructure/files in directory to order by key word
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
			// Split dir names into single names to sort the files
			innerFolders := strings.Split(dir.Name(), ", ")
			log.Println("Splitting folder " + dir.Name() + " into " + strings.Join(innerFolders, " "))
			log.Println("Processing file " + innerFile.Name())

			// General path for deduplicated file collection
			orderedFilesFolder := currentPath + "orderedFiles/"
			generalFolderDestination := orderedFilesFolder + "deduped/"

			// Copy the file into the split folders
			for _, directoryToPutStuffIn := range innerFolders {
				// This is the folder of the file
				pathForFile := orderedFilesFolder + directoryToPutStuffIn + "/"
				// This will be the path for the copied file
				filePath := pathForFile + directoryToPutStuffIn + innerFile.Name()

				if err := os.MkdirAll(pathForFile, 0755); err != nil {
					log.Fatal(err.Error())
				}
				// This is the path of the original file to copy
				pathOfFileToCopy := pathOfDirectory + "/" + innerFile.Name()
				// Execute the copy process
				handleError(copyFileToFolder(pathOfFileToCopy, filePath, generalFolderDestination))

			}

		}

	}

}

func copyFileToFolder(source, destination, genDestination string) error {

	input, err := ioutil.ReadFile(source)
	if err != nil {
		log.Println("Error reading", source)
		return err
	}

	handleError(copyFileToDestination(destination, input))
	// To have one folder with all files de-duplciated
	handleError(copyFileToDestination(genDestination, input))

	return nil
}

func copyFileToDestination(destination string, file []byte) error {
	err = ioutil.WriteFile(destination, file, 0644)
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
