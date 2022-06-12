package util

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var (
	currentPath string
	err         error
)

func CopyFileToFolder(source, destination, genDestination string) error {

	input, err := ioutil.ReadFile(source)
	if err != nil {
		log.Println("Error reading", source)
		return err
	}

	HandleError(CopyFileToDestination(destination, input))
	// To have one folder with all files de-duplicated
	HandleError(CopyFileToDestination(genDestination, input))

	return nil
}

func CopyFileToDestination(destination string, file []byte) error {
	err := ioutil.WriteFile(destination, file, 0644)
	if err != nil {
		log.Println("Error creating", destination)
		return err
	}
	return nil
}

func HandleError(err error) {
	if err != nil {
		log.Fatalln("Program encountered an error " + err.Error())
	}
}

func GetCurrentWorkingPath() (string, error) {
	if len(os.Args) < 2 {
		// Get Current path to work in
		currentPath, err = os.Getwd()
		currentPath = currentPath + "/"
		if err != nil {
			return "", err
		}
	} else {
		currentPath = os.Args[1]
	}

	if currentPath == "" {
		return "", errors.New("current working path returned empty")
	}
	log.Println("Working in current directory: " + currentPath)
	return currentPath, nil
}
