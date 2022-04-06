package main

import (
	_ "embed"
	"fmt"
)

func PrintSmiley() {
	fmt.Println("😀")
}

func PrintUnicodeEmoticons() {
	for r := 0x1F600; r <= 0x1F64F; r++ {
		fmt.Printf("%c ", r)
	}
}

//go:embed emoji-test.txt
var emojiTest string

func PrintEmojiTest() {
	fmt.Println(emojiTest)
}

func main() {
	PrintEmojiTest()
}
