package jellyfinsubtitles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SRTSub struct {
	index     string
	timestamp string
	text      string
}

func addDelay(original string, delay int) string {
	all := strings.Split(original, ":")
	hours, err := strconv.Atoi(strings.ReplaceAll(all[0], " ", ""))
	if err != nil {
		log.Fatal(original, err)
	}
	minutes, err := strconv.Atoi(strings.ReplaceAll(all[1], " ", ""))
	if err != nil {
		log.Fatal(original, err)
	}
	sec := strings.Split(all[2], ",")
	seconds, err := strconv.Atoi(strings.ReplaceAll(sec[0], " ", ""))
	if err != nil {
		log.Fatal(original, err)
	}
	milliseconds, err := strconv.Atoi(strings.ReplaceAll(sec[1], " ", ""))
	if err != nil {
		log.Fatal(original, err)
	}

	msTotal := (hours * 3600 * 1000) + (minutes * 60 * 1000) + (seconds * 1000) + milliseconds + delay
	hours = msTotal / (60 * 60 * 1000)
	msTotal %= (60 * 60 * 1000)

	minutes = msTotal / (60 * 1000)
	msTotal %= (60 * 1000)

	seconds = msTotal / 1000
	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, seconds, milliseconds)
}

func writeToSRT(path string, info []SRTSub) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(path, err, info)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range info {
		newLine := fmt.Sprintf("%v\n%v\n%v\n", line.index, line.timestamp, line.text)
		_, err := writer.WriteString(newLine)
		if err != nil {
			fmt.Println("error while writing to file", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error while flushing", err)
		return
	}
}

// Delay is in milliseconds
func handleSRT(path string, delay int) []SRTSub {
	read, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	content := strings.Split(string(read), "\n")

	var file []SRTSub
	var latest []string
	for _, line := range content {
		line = strings.ReplaceAll(line, "\r", "")
		number, err := strconv.Atoi(line)
		if err == nil {
			if number == 1 {
				continue
			}

			index, timestamp := latest[0], latest[1]
			text := strings.Join(latest[2:], "\n")

			tempSRT := SRTSub{index: index, timestamp: timestamp, text: text}
			file = append(file, tempSRT)
			latest = []string{line}
		}
		if err != nil {
			latest = append(latest, line)
		}
	}

	for index, info := range file {
		timestamps := strings.Split(info.timestamp, "-->")
		start := addDelay(timestamps[0], delay)
		finish := addDelay(timestamps[1], delay)
		timestamp := strings.Join([]string{start, finish}, " --> ")
		file[index].timestamp = timestamp
	}

	return file
}

func ShiftSubtitles() {
	var path string
	fmt.Println("Enter the file path")
	fmt.Scan(&path)

	var dealy int
	fmt.Print("How much delay (in milliseconds)")
	fmt.Scan(&dealy)

	file := handleSRT(path, dealy)
	writeToSRT(path, file)
}
