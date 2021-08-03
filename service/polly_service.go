package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"io"
	"log"
	"os"
)

type PollyService interface {
	Synthesize(text, filename string) error
}

type pollyConfig struct {
	VoiceId string
}

const (
	AUDIO_FORMAT   = "mp3"
	KIMBERLY_VOICE = "Kimberly"
	JOEY_VOICE     = "Joey"
)

func Log(err error) {
	if err != nil {
		log.Printf("%+v", err)
	}
}

func NewKimberlyPollyService() PollyService {
	return &pollyConfig{
		VoiceId: KIMBERLY_VOICE,
	}
}

func NewJoeyPollyService() PollyService {
	return &pollyConfig{
		VoiceId: JOEY_VOICE,
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
		OutputFormat: aws.String(AUDIO_FORMAT),
		Text:         aws.String(text),
		VoiceId:      aws.String(config.VoiceId),
	}

	output, err := pollyClient.SynthesizeSpeech(input)

	Log(err)

	outFile, err := os.Create(filename)

	Log(err)

	defer Log(outFile.Close())

	_, err = io.Copy(outFile, output.AudioStream)

	Log(err)

	return nil
}
