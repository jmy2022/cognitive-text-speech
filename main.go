package main

import (
	"fmt"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/audio"
	"github.com/Microsoft/cognitive-services-speech-sdk-go/speech"
	"time"
)

func main() {

	//file := "Users/jimmy/Music/Untitled-2024-05-01-11-33-14/export/session.mp3"
	//
	//err := ffmpeg.Input("Users/jimmy/Music/Untitled-2024-05-01-11-33-14/export/session.mp3", ffmpeg.KwArgs{"ss": 1}).
	//	Output("./sample_data/out1.mp3", ffmpeg.KwArgs{"t": 1}).OverWriteOutput().Run()
	//log.Println(err) -ac 2 -f wav
	//ffmpeg  -i /Users/jimmy/Desktop/Lily产品设计/20240426原文.mp4 -c:a pcm_s32le -ss 100 -t 10 ./out1.mp4 -y "c": "a pcm_s32le"
	//  -c:a pcm_s32le -ss 100 -t 10 ./out1.mp4 -y
	//err :=

	azaudio()
}

func azaudio() {
	subscription := "be95c2216c1548c49760a88a222f485c"
	region := "eastasia"
	file := "./out1s2s.mp4"

	audioConfig, err := audio.NewAudioConfigFromWavFileInput(file)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer audioConfig.Close()
	config, err := speech.NewSpeechConfigFromSubscription(subscription, region)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer config.Close()
	speechRecognizer, err := speech.NewSpeechRecognizerFromConfig(config, audioConfig)
	if err != nil {
		fmt.Println("Got an error: ", err)
		return
	}
	defer speechRecognizer.Close()
	speechRecognizer.SessionStarted(func(event speech.SessionEventArgs) {
		defer event.Close()
		fmt.Println("Session Started (ID=", event.SessionID, ")")
	})
	speechRecognizer.SessionStopped(func(event speech.SessionEventArgs) {
		defer event.Close()
		fmt.Println("Session Stopped (ID=", event.SessionID, ")")
	})

	task := speechRecognizer.RecognizeOnceAsync()
	var outcome speech.SpeechRecognitionOutcome
	select {
	case outcome = <-task:
	case <-time.After(5 * time.Second):
		fmt.Println("Timed out")
		return
	}
	defer outcome.Close()
	if outcome.Error != nil {
		fmt.Println("Got an error: ", outcome.Error)
	}
	fmt.Println("Got a recognition!")
	fmt.Println(outcome.Result.Text)
}
