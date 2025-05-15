package main

import (
	"fmt"

	"jellyfin/subtitles/jellyfinsubtitles"
)

func main() {
	for {
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
			fmt.Println("See you next time :)")
			return
		case 1:
			jellyfinsubtitles.RenameSubtitles()
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
	}
}
