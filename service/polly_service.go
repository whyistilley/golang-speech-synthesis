package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/whyistilley/golang-speech-synthesis/helpers"
	"io"
	"os"
)

type PollyService interface {
	Synthesize(text, filename string) error
}

type pollyConfig struct {
	VoiceId      string
	OutputFormat string
	TextType     string
	SampleRate   string
}

func NewPollyService(VoiceId, OutputFormat, TextType, SampleRate string) PollyService {
	return &pollyConfig{
		VoiceId:      VoiceId,
		OutputFormat: OutputFormat,
		TextType:     TextType,
		SampleRate:   SampleRate,
	}
}

func createPollyClient() *polly.Polly {
	clientSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return polly.New(clientSession)
}

func (config *pollyConfig) Synthesize(text, filename string) error {
	pollyClient := createPollyClient()

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String(config.OutputFormat),
		Text:         aws.String(text),
		VoiceId:      aws.String(config.VoiceId),
		TextType:     aws.String(config.TextType),
		SampleRate:   aws.String(config.SampleRate),
	}

	output, err := pollyClient.SynthesizeSpeech(input)

	if err != nil {
		return err
	}

	fName := fmt.Sprintf("%s.%s", filename, config.OutputFormat)

	outFile, err := os.Create(fName)

	if err != nil {
		return err
	}

	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {
			helpers.Log(err)
		}
	}(outFile)

	_, err = io.Copy(outFile, output.AudioStream)

	if err != nil {
		return err
	}

	return nil
}
