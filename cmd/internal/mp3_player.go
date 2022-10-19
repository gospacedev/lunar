package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	tb "github.com/nsf/termbox-go"
)

// Plays mp3 file
func MusicPlayer(file string, name string) {
	f, err := os.Open(file)
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
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speaker.Play(volume)

	//initialize termbox
	tbErr := tb.Init()
	if tbErr != nil {
		panic(tbErr)
	}
	defer tb.Close()

	// Print audio file name and key controls
	fmt.Println("Playing " + strings.Replace(name, ".mp3", "", 1) + "...")
	fmt.Println("Audio controls:")
	fmt.Println("Pause and play music: [ENTER]")
	fmt.Println("Volume: [↓ ↑]")
	fmt.Println("Back to menu: [BACKSPACE]")
	fmt.Println("Quit Lunar: [ESC]")
	
	// Detect keys
	for {
		event := tb.PollEvent()
		
		speaker.Lock()

		switch {
		case event.Key == tb.KeyEnter:
			ctrl.Paused = !ctrl.Paused
		case event.Key == tb.KeyArrowUp:
			volume.Volume += 0.2
		case event.Key == tb.KeyArrowDown:
			volume.Volume -= 0.2
		case event.Key == tb.KeyBackspace:
			Start()
		case event.Key == tb.KeyEsc:
			os.Exit(0)
		}

		//max and min volume
		switch {
		case volume.Volume >= 2:
			volume.Volume = 2
		case volume.Volume <= -2:
			volume.Volume = -2
		}

		speaker.Unlock()
	}
}
