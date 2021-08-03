package main

import (
	"fmt"
	"github.com/whyistilley/golang-speech-synthesis/service"
)

var (
	kimberly = service.NewKimberlyPollyService()
	joey     = service.NewJoeyPollyService()
)

func main() {
	err := kimberly.Synthesize("Hi, I am Kimberly. How are you?", "kimberly.mp3")

	if err != nil {
		fmt.Println(err)
	}

	err = joey.Synthesize("Hi, I am Joey. How are you?", "joey.mp3")

	if err != nil {
		fmt.Println(err)
	}
}
