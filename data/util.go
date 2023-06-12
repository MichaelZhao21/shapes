package data

import (
	"os"
)

// MakeDirIfNotExists creates a directory if it does not exist
func MakeDirIfNotExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		GetInfoLogger().Println("Creating directory:", dir)
		err2 := os.Mkdir(dir, os.ModePerm)
		if err2 != nil {
			GetErrorLogger().Println("Error creating directory:", dir)
			os.Exit(1)
		}
	}
}

// GetFileNameList returns a list of file names in a directory
// Does not include the file extension of mp3
func GetFileNameList(dir string) (fileNames []string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		GetErrorLogger().Println("Error reading directory:", dir)
		os.Exit(1)
	}

	for _, file := range files {
		// Remove the .mp3 extension
		fileNames = append(fileNames, file.Name()[0:len(file.Name())-4])
	}

	return fileNames
}

// Contains returns true if the string is in the slice
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
