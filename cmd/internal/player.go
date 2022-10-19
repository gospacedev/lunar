package internal

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	tb "github.com/nsf/termbox-go"
)

// Plays mp3 file
func AudioPlayer(file string, name string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

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
	speedy := beep.ResampleRatio(4, 1, volume)
	speaker.Play(speedy)

	//initialize termbox
	tbErr := tb.Init()
	if tbErr != nil {
		panic(tbErr)
	}
	defer tb.Close()

	// Print audio file name and key controls
	selectedAudio := "Playing " + strings.Replace(name, ".mp3", "", 1)

	p := widgets.NewParagraph()
	p.Title = "Lunar"
	p.Text = selectedAudio
	p.SetRect(0, 0, 40, 3)

	c := widgets.NewParagraph()
	c.Title = "Audio Controls"
	c.Text = `Pause and play music: [ENTER]
Volume: [↓ ↑]
Speed:  [← →]
Noraml Speed: [Ctrl + N]
Back to menu: [BACKSPACE]
Quit Lunar: [ESC]
	`
	c.SetRect(0, 4, 40, 12)

	ui.Render(p, c)

	// Detect keys
	for {
		event := tb.PollEvent()

		speaker.Lock()

		switch {
		case event.Key == tb.KeyEnter: // puase audio
			ctrl.Paused = !ctrl.Paused
		case event.Key == tb.KeyArrowUp: // increase volume
			volume.Volume += 0.2
		case event.Key == tb.KeyArrowDown: // decrease volume
			volume.Volume -= 0.2
		case event.Key == tb.KeyArrowRight: // increase speed by x1.1
			speedy.SetRatio(speedy.Ratio() + 0.1)
		case event.Key == tb.KeyArrowLeft: // decrease speed by x1.1
			speedy.SetRatio(speedy.Ratio() - 0.1)
		case event.Key == tb.KeyCtrlN: // Normalize speed
			speedy.SetRatio(1)
		case event.Key == tb.KeyBackspace: // go back to menu
			Start()
		case event.Key == tb.KeyEsc: // Exit Lunar
			os.Exit(0)
		}

		//maximum and minimum volume and speed
		switch {
		case volume.Volume >= 2:
			volume.Volume = 2
		case volume.Volume <= -2:
			volume.Volume = -2
		case speedy.Ratio() >= 2:
			speedy.SetRatio(2)
		case speedy.Ratio() <= 0.5:
			speedy.SetRatio(0.5)
		}

		speaker.Unlock()
	}
}
