package jellyfinsubtitles

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Subtitles struct {
	original string
	modified string
}

func changeName(videos []string, subtitels []string, videoExt string, subExt string, lang string, episodeIndicatorVideo string, episodeIndicatorSubs string) []Subtitles {
	var changedSubtitles []Subtitles
	for _, video := range videos {
		episodeNumber := getNumberEpisode(video, episodeIndicatorVideo)
		for _, subtitle := range subtitels {
			epNumber := getNumberEpisode(subtitle, episodeIndicatorSubs)

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

func RenameSubtitles() {
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
	var episodeIndicatorVideo string
	var episodeIndicatorSubs string
	var lang string

	fmt.Print("Enter the video extension\n")
	fmt.Scan(&videoExt)
	fmt.Print("Enter the subtitle extension\n")
	fmt.Scan(&subExt)
	fmt.Print("Which lang do you need to add? (ex: EN, ES, IT)\n")
	fmt.Scan(&lang)
	fmt.Print("What is before the episode number of the videos (ex. EP00 write EP)\n")
	fmt.Scan(&episodeIndicatorVideo)
	fmt.Print("What is before the episode number of the subtitles (ex. EP00 write EP)\n")
	fmt.Scan(&episodeIndicatorSubs)

	videos := listFiles(dir, videoExt)
	subs := listFiles(dir, subExt)

	correctSubs := changeName(videos, subs, videoExt, subExt, lang, episodeIndicatorVideo, episodeIndicatorSubs)

	renameFiles(correctSubs)
	fmt.Print("Everything should be done!\n")
}
