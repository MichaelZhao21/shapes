package data

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strings"
)

type Song struct {
	Name    string
	Artist  string
	Url     string // Direct download link
	PageUrl string // Link to the song's page
}

func LoadSongs() (songs map[string]Song) {
	dataDir := GetDataDir()

	// Get name of all files of the approved downloaded songs
	// and store them in a slice
	songList := GetFileNameList(dataDir + "/songs/approved")

	// Load list of all songs
	rawSongFile := ReadApprovedSongs()
	songs = ParseApprovedSongs(rawSongFile)

	// Download all songs that are not already downloaded
	for k := range songs {
		if !Contains(songList, k) {
			// Download song
			DownloadSong(songs[k])
		}
	}

	return songs
}

// ParseApprovedSongs parses the approved_songs file
// and returns a "set" of all approved songs
func ParseApprovedSongs(rawSongFile string) (songs map[string]Song) {
	songs = make(map[string]Song)

	r := csv.NewReader(strings.NewReader(rawSongFile))

	for {
		record, err := r.Read()
		if err != nil {
			break
		}

		songs[record[0]] = Song{
			Name:    strings.Trim(record[0], " "),
			Artist:  strings.Trim(record[1], " "),
			Url:     strings.Trim(record[2], " "),
			PageUrl: strings.Trim(record[3], " "),
		}
	}

	return songs
}

// DownloadSong downloads a song from the internet
// based on the song struct's url field
func DownloadSong(song Song) {
	// TODO: Put this in a goroutine
	GetInfoLogger().Println("Downloading song:", song.Name, "from", song.Url)

	// Create the file to store the downloaded song
	fileName := GetDataDir() + "/songs/approved/" + song.Name + ".mp3"
	file, err := os.Create(fileName)
	if err != nil {
		GetErrorLogger().Println("Could not create song file:", fileName)
		os.Exit(1)
	}
	defer file.Close()

	// Get the data from the url
	res, err := http.Get(song.Url)
	if err != nil {
		GetErrorLogger().Println("Could not download song:", song.Name, "from", song.Url, "|", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	// Make sure request was successful
	if res.StatusCode != http.StatusOK {
		GetErrorLogger().Println("Error when downloading song:", song.Name, "from", song.Url, "|", res.Status)
		os.Exit(1)
	}

	// Write the body to file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		GetErrorLogger().Println("Could not write song to file:", fileName)
		os.Exit(1)
	}

	GetInfoLogger().Println("Download complete:", song.Name)
}

func OpenSongFile(song Song) (file *os.File) {
	fileName := GetDataDir() + "/songs/approved/" + song.Name + ".mp3"
	file, err := os.Open(fileName)
	if err != nil {
		GetErrorLogger().Println("Could not open song file:", fileName)
		os.Exit(1)
	}

	return file
}
