package main

import (
	"github.com/whyistilley/golang-speech-synthesis/helpers"
	"github.com/whyistilley/golang-speech-synthesis/service"
	"io/ioutil"
)

const (
	// VoiceId voices to pick from = { "Amy", "Matthew"}
	VoiceId      = "Matthew"
	OutputFormat = "mp3"
	TextType     = "ssml"
	SampleRate   = "24000"
)

var (
	polly = service.NewPollyService(VoiceId, OutputFormat, TextType, SampleRate)
)

func main() {
	text, err := ioutil.ReadFile("ssml_script.ssml")

	helpers.Log(err)

	err = polly.Synthesize(string(text), VoiceId)

	helpers.Log(err)
}
