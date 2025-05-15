package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"jellyfin/subtitles/jellyfinsubtitles"
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

func getNumberEpisode(path string, episodeIndicator string) int {
	pattern := fmt.Sprintf(`[%s][%s]\d{1,99}`, episodeIndicator, episodeIndicator)

	fmt.Println(pattern)

	re := regexp.MustCompile(pattern)
	match := re.FindString(path)
	match = strings.ReplaceAll(match, episodeIndicator, "")
	match = strings.ReplaceAll(match, episodeIndicator, "")
	if match == "" {
		return -1
	}
	index, err := strconv.Atoi(match)
	if err != nil {
		return -1
	}
	return index
}

func changeName(videos []string, subtitels []string, videoExt string, subExt string, lang string, episodeIndicator string) []Subtitles {
	var changedSubtitles []Subtitles
	for _, video := range videos {
		episodeNumber := getNumberEpisode(video, episodeIndicator)
		for _, subtitle := range subtitels {
			epNumber := getNumberEpisode(subtitle, episodeIndicator)
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
		fmt.Println(sub)

		err := os.Rename(sub.original, sub.modified)
		if err != nil {
			fmt.Print("Error with", err, "\n")
		}
	}
}

func main() {
	var choice int
	for {
		fmt.Printf("\n\nWhat would you like to do?\n1) Rename subtitles\n2) Shift subtitles\n3) Calculate time\n4) Concat videos\n0) Close\n")
		fmt.Scanf("%d", &choice)

		fmt.Println("you chose ", choice)

		if choice >= 0 && choice <= 4 {
			break
		}
	}

	switch choice {
	case 0:
		return
	case 1:
		fmt.Println("Subtitles")
		break
	case 2:
		jellyfinsubtitles.ShiftSubtitles()
		break
	case 3:
		jellyfinsubtitles.Calculus()
		return
	case 4:
		jellyfinsubtitles.ConcatVideos()
		break
	default:
		return
	}

	reader := bufio.NewReader(os.Stdin)

	var dir string
	for {
		fmt.Print("Enter the folder path needs to be absolute path\n")
		dir, _ = reader.ReadString('\n')
		dir = strings.ReplaceAll(dir, "\n", "")
		dir = filepath.Join(dir)

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println("Please insert a valid path")
		} else {
			break
		}
	}

	var videoExt string
	var subExt string
	var episodeIndicator string
	var lang string

	fmt.Print("Enter the video extension\n")
	fmt.Scan(&videoExt)
	fmt.Print("Enter the subtitle extension\n")
	fmt.Scan(&subExt)
	fmt.Print("What is before the episode number (ex. EP00 write EP)\n")
	fmt.Scan(&episodeIndicator)
	fmt.Print("Which lang do you need to add? (ex: EN, ES, IT)\n")
	fmt.Scan(&lang)

	videos := listFiles(dir, videoExt)
	subs := listFiles(dir, subExt)

	correctSubs := changeName(videos, subs, videoExt, subExt, lang, episodeIndicator)

	renameFiles(correctSubs)
	fmt.Print("Everything should be done!\n")
}
