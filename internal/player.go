package internal

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// GenrericDecoder detects different audio formats then decodes it
func GenrericDecoder(name string, f *os.File) (beep.StreamSeekCloser, beep.Format, error) {
	switch {
		case strings.HasSuffix(name, ".mp3"):
			return mp3.Decode(f)
		case strings.HasSuffix(name, ".wav"):
			return wav.Decode(f)
		case strings.HasSuffix(name, ".flac"):
			return flac.Decode(f)
		case strings.HasSuffix(name, ".ogg"):
			return vorbis.Decode(f)
	}

	// the deafault decoder is mp3
	return mp3.Decode(f)
}

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

	streamer, format, err := GenrericDecoder(name, f)
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

	p := widgets.NewParagraph()
	p.Title = "Playing"
	p.Text = name
	p.SetRect(0, 0, 40, 3)
	p.BorderStyle.Fg = ui.ColorCyan
	p.TitleStyle.Fg = ui.ColorYellow

	volGauge := widgets.NewGauge()
	volGauge.Title = "Volume"
	volGauge.Percent = 50
	volGauge.SetRect(0, 3, 40, 6)
	volGauge.BorderStyle.Fg = ui.ColorCyan
	volGauge.TitleStyle.Fg = ui.ColorYellow

	speedGauge := widgets.NewGauge()
	speedGauge.Title = "Speed"
	speedGauge.Percent = 50
	speedGauge.SetRect(0, 6, 40, 9)
	speedGauge.BorderStyle.Fg = ui.ColorCyan
	speedGauge.TitleStyle.Fg = ui.ColorYellow

	c := widgets.NewParagraph()
	c.Text = `Pause / Play: [ENTER]
Volume: [↓ ↑]
Speed:  [← →]
Normal Speed: [N]
Back to Menu: [Backspace]
	`
	c.SetRect(0, 9, 40, 16)
	c.BorderStyle.Fg = ui.ColorCyan

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	draw := func() {
		ui.Render(p, c, volGauge, speedGauge)
	}

	// Detect keys
	for {

		select {
		case e := <-uiEvents:
			switch e.ID {
			case "<Enter>": // pause audio
				ctrl.Paused = !ctrl.Paused
			case "<Up>": // increase volume
				volume.Volume += 0.1
				volGauge.Percent += 2
			case "<Down>": // decrease volume
				volume.Volume -= 0.1
				volGauge.Percent -= 2
			case "<Right>": // increase speed by x1.1
				speedy.SetRatio(speedy.Ratio() + 0.1)
				speedGauge.Percent += 2
			case "<Left>": // decrease speed by x1.1
				speedy.SetRatio(speedy.Ratio() - 0.1)
				speedGauge.Percent -= 2
			case "n": // Normalize speed
				speedy.SetRatio(1)
				speedGauge.Percent = 50
			case "<C-<Backspace>>": // go back to menu
				ctrl.Paused = !ctrl.Paused
				Menu()
			}
		case <-ticker:
			draw()
		}

		//set max and min volume and speed
		switch {
		case volume.Volume >= 2:
			volume.Volume = 2
			volGauge.Percent = 100
		case volume.Volume <= -2:
			volume.Volume = -2
			volGauge.Percent = 0
		case speedy.Ratio() >= 2:
			speedy.SetRatio(2)
			speedGauge.Percent = 100
		case speedy.Ratio() <= 0.5:
			speedy.SetRatio(0.5)
			speedGauge.Percent = 0
		}
	}
}
