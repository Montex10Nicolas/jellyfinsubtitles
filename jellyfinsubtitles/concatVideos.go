package jellyfinsubtitles

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
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
	pattern := fmt.Sprintf(`[%s][%s]\d{1,99}`, episodeIndicator, episodeIndicator)
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

func number2Digs(n int) string {
	sn := strconv.Itoa(n)
	if n > 10 {
		return sn
	} else {
		return "0" + sn
	}
}

func ConcatVideos() {
	var inputDir string
	var outputDir string
	var videoExt string
	var episodeChar string
	var episodeName string

	var splitChar string

	var confirm string = "Y"

	reader := bufio.NewReader(os.Stdin)
	for {
		var dirExists = false
		for {
			fmt.Println("\n\nEnter the location of the videos (absolute path)")
			inputDir, _ = reader.ReadString('\n')
			inputDir = strings.ReplaceAll(inputDir, "\n", "")
			inputDir = filepath.Join(inputDir)

			fmt.Println(inputDir)

			if _, err := os.Stat(inputDir); os.IsNotExist(err) {
				fmt.Println("Please choose an existsing directory")
			} else {
				dirExists = true
			}
			if dirExists {
				break
			}
		}

		fmt.Println("Where to output the files (absolute path)")
		outputDir, _ = reader.ReadString('\n')
		outputDir = filepath.Join(outputDir)

		if outputDir[len(outputDir)-1:] != "/" {
			outputDir += "/"
		}

		fmt.Println("Which extension is the video (es. .mp4 .mkv)")
		fmt.Scan(&videoExt)
		fmt.Println("How which are the letters that preceed the episode (ex. EP00 E00)")
		fmt.Scan(&episodeChar)
		fmt.Println("How the split are indicated (ex. use ')' if '(1/3)' is used")
		fmt.Scan(&splitChar)
		fmt.Println("How to name every video, this will be followed by the episode indicator specified before and the number")
		episodeName, _ = reader.ReadString('\n')
		episodeName = strings.ReplaceAll(episodeName, "\n", "")

		fmt.Printf("Recap:\nInput Path: %s\nOutput Path: %s\nVideo extension: %s\nEpisode indicator: %s\nSplit: %s\nEpisode name: %s\n", inputDir, outputDir, videoExt, episodeChar, splitChar, episodeName)

		fmt.Println("Is this correct? (Y/n)")
		fmt.Scan(&confirm)

		if strings.ToLower(confirm) == "y" {
			break
		} else {
			fmt.Println("Let's try again")
		}
	}

	videos := listFiles(inputDir, videoExt)
	var m map[int][]string = make(map[int][]string)

	for _, v := range videos {
		n := getNumberEpisode(v, episodeChar)
		m[n] = append(m[n], v)
	}

	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	for episodeNumber, value := range m {
		fileToConcat := "/tmp/file.txt"
		file, err := os.Create(fileToConcat)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		writer := bufio.NewWriter(file)

		end := strings.Split(value[0], splitChar)[1]
		outfile := fmt.Sprintf("%s%s %s%s", outputDir, episodeName, number2Digs(episodeNumber), end)

		for _, filename := range value {
			_, err = writer.WriteString("file " + "'" + filename + "'" + "\n")
			if err != nil {
				panic(err)
			}
		}
		writer.Flush()

		fmt.Printf("Doing %s %d output in: %s\n", episodeName, episodeNumber, outfile)
		cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", "/tmp/file.txt", "-c", "copy", outfile)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing command: %v\n", err)
			fmt.Printf("Command output: %s\n", output)
			panic(err)
		}
		fmt.Printf("Done %s %d in %s\n\n", episodeName, episodeNumber, outfile)
	}

	fmt.Printf("Everything is done :)")
}
