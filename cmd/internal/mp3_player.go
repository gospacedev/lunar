package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	tb "github.com/nsf/termbox-go"
)

// Plays mp3 file
func MusicPlayer(music string, name string) {
	f, err := os.Open(music)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	speaker.Play(ctrl)

	//initialize termbox
	tbErr := tb.Init()
	if tbErr != nil {
		panic(tbErr)
	}
	defer tb.Close()

	// Print audio file name and key controls
	fmt.Println("Playing " + strings.Replace(name, ".mp3", "", 1))
	fmt.Println("Pause and play: [ENTER]")
	fmt.Println("Back: [BACKSPACE]")
	fmt.Println("Quit: [ESC]")
	
	// Detect keys
	for {
		event := tb.PollEvent()

		switch {
		case event.Key == tb.KeyEnter:
			ctrl.Paused = !ctrl.Paused
		case event.Key == tb.KeyBackspace:
			Start()
		case event.Key == tb.KeyEsc:
			os.Exit(0)
		}
	}
}
