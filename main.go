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

type Subtitles struct {
	original string
	modified string
}

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
	pattern := `[Ee]\d{1,99}`
	re := regexp.MustCompile(pattern)
	match := re.FindString(path)
	match = strings.ReplaceAll(match, "E", "")
	match = strings.ReplaceAll(match, "e", "")
	if match == "" {
		return -1
	}
	index, err := strconv.Atoi(match)
	if err != nil {
		return -1
	}
	return index
}

func changeName(videos []string, subtitels []string, videoExt string, subExt string, lang string) []Subtitles {
	var changedSubtitles []Subtitles
	for _, video := range videos {
		episodeNumber := getNumberEpisode(video)
		for _, subtitle := range subtitels {
			epNumber := getNumberEpisode(subtitle)
			if epNumber != -1 && epNumber == episodeNumber {
				newSub := strings.ReplaceAll(video, videoExt, lang+"."+subExt)
				changedSubtitles = append(changedSubtitles, Subtitles{original: subtitle, modified: newSub})
			}
		}
	}

	return changedSubtitles
}

func renameFiles(subs []Subtitles) {
	for _, sub := range subs {
		err := os.Rename(sub.original, sub.modified)
		if err != nil {
			fmt.Print("Error with", err, "\n")
		}
	}
}

func main() {
	var dir string
	fmt.Print("Enter the folder path needs to be absolute path\n")
	fmt.Scan(&dir)

	var videoExt string
	var subExt string
	var lang string

	fmt.Print("Enter the video extension\n")
	fmt.Scan(&videoExt)
	fmt.Print("Enter the subtitle extension\n")
	fmt.Scan(&subExt)
	fmt.Print("Which lang do you need to add? (ex: EN, ES, IT)\n")
	fmt.Scan(&lang)

	videos := listFiles(dir, videoExt)
	subs := listFiles(dir, subExt)

	correctSubs := changeName(videos, subs, videoExt, subExt, lang)

	renameFiles(correctSubs)
	fmt.Print("Everything should be done!\n")
}
