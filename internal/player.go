/*
Copyright © 2022 Grantley Cullar <grantcullar@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package internal

import (
	"fmt"
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
	"github.com/spf13/viper"

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

// Play and control an audio file
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

	vp := viper.New()

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(home)

	err1 := vp.ReadInConfig()
	if err1 != nil {
		fmt.Println("Error: Cannot read config file")
	}

	TitleThemeColor := vp.GetInt("titlethemecolor")
	BorderThemeColor := vp.GetInt("borderthemecolor")

	if TitleThemeColor == 0 || BorderThemeColor == 0 {
		TitleThemeColor = 3
		BorderThemeColor = 6
	}

	p := widgets.NewParagraph()
	p.Title = "Playing"
	p.Text = name
	p.SetRect(0, 0, 50, 3)
	p.TitleStyle.Fg = ui.Color(TitleThemeColor)
	p.BorderStyle.Fg = ui.Color(BorderThemeColor)

	posGauge := widgets.NewGauge()
	posGauge.Title = "Position"
	posGauge.Percent = 0
	posGauge.SetRect(0, 3, 50, 6)
	posGauge.TitleStyle.Fg = ui.Color(TitleThemeColor)
	posGauge.BorderStyle.Fg = ui.Color(BorderThemeColor)

	volGauge := widgets.NewGauge()
	volGauge.Title = "Volume"
	volGauge.Percent = 50
	volGauge.SetRect(0, 6, 50, 9)
	volGauge.TitleStyle.Fg = ui.Color(TitleThemeColor)
	volGauge.BorderStyle.Fg = ui.Color(BorderThemeColor)

	speedGauge := widgets.NewGauge()
	speedGauge.Title = "Speed"
	speedGauge.Percent = 50
	speedGauge.SetRect(0, 9, 50, 12)
	speedGauge.TitleStyle.Fg = ui.Color(TitleThemeColor)
	speedGauge.BorderStyle.Fg = ui.Color(BorderThemeColor)

	c := widgets.NewParagraph()
	c.Text = `Pause / Play: [ENTER]
Position: [← →]
Volume: [↓ ↑]
Speed:  [A / S]
Normal Speed: [N]
Back to Menu: [BACKSPACE]
	`
	c.SetRect(0, 12, 50, 19)
	c.BorderStyle.Fg = ui.Color(BorderThemeColor)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C

	draw := func() {
		ui.Render(p, c, posGauge, volGauge, speedGauge)
	}

	// Detect keys
	for {
		newPos := streamer.Position()

		position := format.SampleRate.D(streamer.Position())
		length := format.SampleRate.D(streamer.Len())

		positionStatus := 100 * (float32(position.Round(time.Second)) / float32(length.Round(time.Second)))
		posGauge.Percent = int(positionStatus)

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

			case "<Left>", "<Right>":
				speaker.Lock()
				if e.ID == "<Left>" {
					newPos -= int(format.SampleRate.N(time.Second))
				} else if e.ID == "<Right>" {
					newPos += int(format.SampleRate.N(time.Second))
				}
				if newPos < 0 {
					newPos = 0
				}
				if newPos >= streamer.Len() {
					newPos = streamer.Len() - 1
				}
				if err := streamer.Seek(newPos); err != nil {
					fmt.Println(err)
				}
				speaker.Unlock()

			case "a":
				speedy.SetRatio(speedy.Ratio() - 0.1)
				speedGauge.Percent -= 2
			
			case "s":
				speedy.SetRatio(speedy.Ratio() + 0.1)
				speedGauge.Percent += 2

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
