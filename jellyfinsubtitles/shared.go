package jellyfinsubtitles

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

	filetype = strings.ReplaceAll(filetype, ".", "")

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
	var regexPattern string = ""
	for range episodeIndicator {
		regexPattern = fmt.Sprintf("%s[%s]", regexPattern, episodeIndicator)
	}

	pattern := fmt.Sprintf(`%s\d{1,99}`, regexPattern)

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
