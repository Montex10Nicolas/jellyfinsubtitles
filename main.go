package main

import (
	"fmt"

	"jellyfin/subtitles/jellyfinsubtitles"
)

func main() {
	for {
		var choice int
		for {
			options := []string{
				"What would you like to do?",
				"1) Rename subtitles",
				"2) Shift subtitles",
				"3) Calculate time",
				"4) Concat videos",
				"5) Rename Videos (This will change the name of the file to what you want followed by SXXEPXX)",
				"6) Get Lenghts of all videos in a file, downstream",
				"0) Close",
			}

			for _, value := range options {
				fmt.Println(value)
			}

			_, err := fmt.Scanf("%d", &choice)
			if err != nil {
				return
			}

			fmt.Println("you chose ", choice)

			if choice >= 0 && choice <= 6 {
				break
			}
		}

		switch choice {
		case 0:
			fmt.Println("See you next time :)")
		case 1:
			jellyfinsubtitles.RenameSubtitles()
		case 2:
			jellyfinsubtitles.ShiftSubtitles()
		case 3:
			jellyfinsubtitles.Calculus()
		case 4:
			jellyfinsubtitles.ConcatVideos()
		case 5:
			jellyfinsubtitles.RenameVideos()
		case 6:
			jellyfinsubtitles.LengthyInit()
		default:
			return
		}
	}
}
