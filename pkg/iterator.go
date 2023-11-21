package pkg

import (
	"GIN/utils"
)

func Files(createChannel chan utils.FileOperation, data string) {
	counts := utils.FileOperation{}
	for _, char := range data {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			counts.Vowel++
		case '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ':', ';', '\'', '<', '>', ',', '.', '/', '?':
			counts.Punctuation++
		case '\n':
			counts.Nextline++
		case ' ':
			counts.Spaces++
		}
		counts.Chars++
	}
	createChannel <- counts
}
