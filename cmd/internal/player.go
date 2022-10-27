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

	// Print audio file name and key controls
	selectedAudio := strings.Replace(name, ".mp3", "", 1)

	p := widgets.NewParagraph()
	p.Title = "Playing"
	p.Text = selectedAudio
	p.SetRect(0, 0, 40, 3)
	p.BorderStyle.Fg = ui.ColorCyan

	c := widgets.NewParagraph()
	c.Text = `Pause / Play: [ENTER]
Volume: [↓ ↑]
Speed:  [← →]
Normal Speed: [N]
Back to Menu: [Backspace]
Quit Lunar: [Q]
	`

	c.SetRect(0, 4, 40, 12)
	c.TitleStyle.Fg = ui.ColorYellow
	c.BorderStyle.Fg = ui.ColorCyan

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	draw := func() {
		ui.Render(p, c)
	}

	// Detect keys
	for {

		select {
		case e := <-uiEvents:
			switch e.ID {
			case "<Enter>": // pause audio
				ctrl.Paused = !ctrl.Paused
			case "<Up>": // increase volume
				volume.Volume += 0.2
			case "<Down>": // decrease volume
				volume.Volume -= 0.2
			case "<Right>": // increase speed by x1.1
				speedy.SetRatio(speedy.Ratio() + 0.1)
			case "<Left>": // decrease speed by x1.1
				speedy.SetRatio(speedy.Ratio() - 0.1)
			case "n": // Normalize speed
				speedy.SetRatio(1)
			case "<C-<Backspace>>": // go back to menu
				ctrl.Paused = !ctrl.Paused
				Menu()
			case "q": // quit Lunar
				return
			}
		case <-ticker:
			draw()
		}

		//set max and min volume and speed
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
	}
}
