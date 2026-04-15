package jellyfinsubtitles

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func renameVideo(original string, newName string, season int, episode int) {
}

func RenameVideos() {
	var dir string
	var files []os.DirEntry
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter the folder path needs to be absolute path\n")
		dir, _ = reader.ReadString('\n')
		dir = strings.ReplaceAll(dir, "\n", "")
		dir = filepath.Join(dir)

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println("Please insert a valid path", dir)
		} else {
			break
		}
	}

	var season string
	fmt.Printf("What is before the season number: ")
	fmt.Scan(&season)

	var episode string
	fmt.Printf("What is before the episode number: ")
	fmt.Scan(&episode)

	var newEpisodeNmae string
	fmt.Printf("Insert new name")
	fmt.Scan(&newEpisodeNmae)

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Errorf("error %d", err)
	}

	for i := range files {
		fmt.Printf("%d\n", files[i])
	}
}
