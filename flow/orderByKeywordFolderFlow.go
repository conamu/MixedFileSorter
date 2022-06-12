package flow

/*
var (
	currentPath string
	err         error
)

*/

/*
	if len(os.Args) < 2 {
		// Get Current path to work in
		currentPath, err = os.Getwd()
		currentPath = currentPath + "/"
		util.HandleError(err)
	} else {
		currentPath = os.Args[1]
	}
	log.Println("Working in current directory: " + currentPath)

	// Read folderstructure/files in directory to order by key word
	folders, err := os.ReadDir(currentPath)
	util.HandleError(err)

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
		util.HandleError(err)

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
			if err := os.MkdirAll(generalFolderDestination, 0755); err != nil {
				log.Fatal(err.Error())
			}

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
				genPathOfFileToCopy := generalFolderDestination + "/" + innerFile.Name()
				// Execute the copy process
				util.HandleError(util.CopyFileToFolder(pathOfFileToCopy, filePath, genPathOfFileToCopy))

			}

		}

	}

*/
