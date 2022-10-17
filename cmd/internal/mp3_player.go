package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	tb "github.com/nsf/termbox-go"
)

// Plays mp3 file
func MusicPlayer(music string) {
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
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speaker.Play(ctrl)

	tbErr := tb.Init()
    if tbErr != nil {
        panic(tbErr)
    }
    defer tb.Close()

	fmt.Println("Ues enter key to pause and resume: Enter")
	fmt.Println("Press the arrow keys to change volume: ↓ ↑")
	fmt.Println("Press escape key to exit Lunar: Esc")

	// Detect keys
	for {
		event := tb.PollEvent()

		switch {
		case event.Key == tb.KeyEnter:
			ctrl.Paused = !ctrl.Paused
		case event.Key == tb.KeyArrowUp:
			volume.Volume += 0.5
		case event.Key == tb.KeyArrowDown:
			volume.Volume -= 0.5
		case event.Key == tb.KeyEsc:
			os.Exit(0)
		}
	}
}