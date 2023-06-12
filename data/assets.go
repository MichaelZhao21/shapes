package data

import (
	"embed"
	"os"
	"strings"
)

var (
	AssetsFS embed.FS
	DataDir  = "gamedata"
)

// GetGameDir returns the game data directory
// If the game data directory does not exist, it will be created
func GetDataDir() string {
	// Set the data dir based on the relative path of the executable
	// if unset
	if DataDir == "gamedata" {
		cwd, err := os.Getwd()
		if err != nil {
			GetErrorLogger().Println("Error getting current working directory. Please check your permissions.")
			os.Exit(1)
		}
		DataDir = cwd + "/" + DataDir

	}
	// Load the data_dirs file
	dirs := GetDataDirList()

	// Make sure all data directories exist
	for _, dir := range dirs {
		MakeDirIfNotExists(dir)
	}

	return DataDir
}

// Returns a slice of all data directories.
// Each directory is a subdirectory of the game data directory.
// Reads the values from the data_dirs file
func GetDataDirList() (dirs []string) {
	file, err := AssetsFS.ReadFile("assets/data_dirs")
	if err != nil {
		GetErrorLogger().Println("data_dirs file not found, please check for valid executable")
		os.Exit(1)
	}

	// Split the file into a slice of directories (with one directory per line)
	rawDirs := string(file)
	splitDirs := strings.Split(rawDirs, "\n")

	// Append each directory to the dirs slice, prefixing with the data dir
	dirs = append(dirs, DataDir)
	for _, dir := range splitDirs {
		if dir != "" {
			dirs = append(dirs, DataDir+"/"+dir)
		}
	}
	return dirs
}

func ReadApprovedSongs() string {
	// Load the approved songs file
	file, err := AssetsFS.ReadFile("assets/approved_songs.csv")
	if err != nil {
		GetErrorLogger().Println("approved_songs.csv file not found, please check for valid executable")
		os.Exit(1)
	}

	return string(file)
}
