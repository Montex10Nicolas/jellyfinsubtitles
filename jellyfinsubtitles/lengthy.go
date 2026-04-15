package jellyfinsubtitles

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type FFProbeOutput struct {
	Format struct {
		Duration string `json:"duration"`
	} `json:"format"`
}

func secondsToMM(s int64) string {
	d := time.Duration(s) * time.Second
	m := int(d.Minutes())

	return fmt.Sprintf("%02dM", m)
}

func LengthyInit() {
	var inputDir string

	fmt.Println("Lengthy")

	reader := bufio.NewReader(os.Stdin)
	for {

		for {
			dirExists := false
			fmt.Println("\n\nEnter the location of the videos (absolute path)")
			inputDir, _ = reader.ReadString('\n')
			inputDir = strings.ReplaceAll(inputDir, "\n", "")
			inputDir = filepath.Join(inputDir)

			if _, err := os.Stat(inputDir); os.IsNotExist(err) {
				fmt.Println("Please choose an existsing directory")
			} else {
				dirExists = true
			}
			if dirExists {
				break
			}
		}

		// Change working directory
		err := os.Chdir(inputDir)
		if err != nil {
			fmt.Print(err)
		}

		entries, err := os.ReadDir(inputDir)
		if err != nil {
			return
		}

		var directories []os.DirEntry
		mapEntries := make(map[string][]os.DirEntry)

		for _, v := range entries {
			if v.IsDir() {
				directories = append(directories, v)
			} else {
				x := mapEntries[inputDir]
				x = append(x, v)
				mapEntries[inputDir] = x
			}
		}

		for _, entries := range directories {
			currentPath, err := os.ReadDir(entries.Name())
			if err != nil {
				fmt.Println("Something wrong", err)
			}
			for _, entry := range currentPath {
				if entry.IsDir() {
					directories = append(directories, entry)
				} else {
					x := mapEntries[entries.Name()]
					x = append(x, entry)
					mapEntries[entries.Name()] = x
				}
			}
		}

		for base, entries := range mapEntries {
			fmt.Println("Season: ", base)
			for _, file := range entries {
				info, err := file.Info()
				if err != nil {
					fmt.Println(err, info)
				}
				filePath := fmt.Sprintf("%s/%s", base, info.Name())

				cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "format=duration", "-of", "json", filePath)
				out, err := cmd.Output()
				if err != nil {
					log.Fatal(err)
				}

				var res FFProbeOutput
				if err := json.Unmarshal(out, &res); err != nil {
					log.Fatal(err)
				}
				durationInSeconds, err := strconv.ParseFloat(res.Format.Duration, 64)
				if err != nil {
					fmt.Println("Something wrong with convertion", err, res.Format.Duration)
				}

				fmt.Printf("%s: %s\n", info.Name(), secondsToMM(int64(durationInSeconds)))
			}
			fmt.Printf("End of season %s\n\n", base)
		}
	}
}
