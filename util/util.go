package util

import (
	"io/ioutil"
	"log"
)

func CopyFileToFolder(source, destination, genDestination string) error {

	input, err := ioutil.ReadFile(source)
	if err != nil {
		log.Println("Error reading", source)
		return err
	}

	HandleError(CopyFileToDestination(destination, input))
	// To have one folder with all files de-duplciated
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
