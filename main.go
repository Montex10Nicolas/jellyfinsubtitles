package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func listFiles(dir string, filetype string) []string {
	root := os.DirFS(dir)

	files, err := fs.Glob(root, "*."+filetype)

	if err != nil {
		log.Fatal(err)
	}

	var retFiles []string
	for _, v := range files {
		retFiles = append(retFiles, path.Join(dir, v))
	}
	return retFiles
}

func getNumberEpisode(path string) int {
	pattern := `E\d{1,5}`
	re := regexp.MustCompile(pattern)
	match := re.FindString(path)
	match = strings.ReplaceAll(match, "E", "")
	if match == "" {
		return -1
	}
	index, err := strconv.Atoi(match)
	if err != nil {
		return -1
	}
	return index
}

func changeName(videos []string, subtitels []string, videoExt string, subExt string) []string {
	var changedSubtitles []string
	for _, video := range videos {
		episodeNumber := getNumberEpisode(video)
		for _, subtitle := range subtitels {
			epNumber := getNumberEpisode(subtitle)
			if epNumber != -1 && epNumber == episodeNumber {
				newSub := strings.ReplaceAll(video, videoExt, subExt)
				changedSubtitles = append(changedSubtitles, newSub)
			}
		}
	}

	return changedSubtitles
}

func renameFiles(originalNames []string, correctedNames []string) {
	for _, original := range originalNames {
		ogNumber := getNumberEpisode(original)
		for _, newName := range correctedNames {
			newNumber := getNumberEpisode(newName)
			if newNumber != -1 && newNumber == ogNumber {
				err := os.Rename(original, newName)
				if err != nil {
					fmt.Println("There has been an error with ", original, "and", newName)
				}
			}
		}
	}
}

func main() {
	var dir string
	fmt.Print("Enter the folder path\n")
	fmt.Scan(&dir)
	var videoExt string
	var subExt string

	fmt.Print("Enter the video extension\n")
	fmt.Scan(&videoExt)
	fmt.Print("Enter the subtitle extension\n")
	fmt.Scan(&subExt)

	videos := listFiles(dir, videoExt)
	subs := listFiles(dir, subExt)

	correctSubs := changeName(videos, subs, videoExt, subExt)

	renameFiles(subs, correctSubs)
	fmt.Print("Everything should be done!")
}
