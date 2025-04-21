package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
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

func getNumberEpisode(path string) {

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

	fmt.Print("This are the videos\n")
	for _, v := range videos {
		fmt.Println(v)
	}
	fmt.Print("This are the subtitles\n")
	for _, v := range subs {
		fmt.Println(v)
	}
}
