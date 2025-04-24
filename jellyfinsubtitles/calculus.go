package jellyfinsubtitles

import (
	"fmt"
	"math"
)

func seedingTime(size int) string {
	var timeMinutes int
	if size < 50 {
		duration := size * 2
		timeMinutes = 259200 + duration*3600
	} else {
		duration := 100*math.Log(float64(size)) - 219.2023
		timeMinutes = int(duration) * 3600
	}

	fmt.Println(timeMinutes)

	days := timeMinutes / (3600 * 24)
	timeMinutes %= (3600 * 24)

	hours := timeMinutes / (60 * 60)
	timeMinutes %= (60 * 60)

	minutes := timeMinutes / (60)
	timeMinutes %= (60)

	seconds := timeMinutes / 1000
	return fmt.Sprintf("%02d:%02d:%02d:%02d", days, hours, minutes, seconds)
}

func Calculus() {
	var size int
	fmt.Printf("How large is the download\n")
	fmt.Scanf("%d", &size)

	time := seedingTime(size)
	fmt.Printf("You need: %v\n", time)
}
