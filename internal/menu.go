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

	"github.com/spf13/viper"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// The menu where the user can choose what audio file to play
func Menu() {
	vp := viper.New()

	// get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(home)

	// Reading config file
	err1 := vp.ReadInConfig()
	if err1 != nil {
		fmt.Println("Error: Cannot read config file")
	}

	path, err := os.ReadDir(vp.GetString("path"))
	if err != nil {
		fmt.Println("No filepath detected, to add new filepath run: lunar newpath")
		os.Exit(1)
	}

	// get widget theme
	TitleThemeColor := vp.GetInt("titlethemecolor")
	BorderThemeColor := vp.GetInt("borderthemecolor")
	SelectedRowThemeColor := vp.GetInt("SelectedRowThemeColor")

	if TitleThemeColor == 0 || BorderThemeColor == 0 {
		TitleThemeColor = 3
		BorderThemeColor = 6
		SelectedRowThemeColor = 3
	}

	var items []string

	// Get file names in folder
	for _, f := range path {
		if isAudioFile(f.Name()) {
			items = append(items, f.Name())
		}
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "Music Menu"
	l.Rows = items
	l.SelectedRowStyle = ui.NewStyle(ui.Color(SelectedRowThemeColor))
	l.WrapText = false
	l.SetRect(0, 0, 40, len(items)+2)
	l.TitleStyle.Fg = ui.Color(TitleThemeColor)
	l.BorderStyle.Fg = ui.Color(BorderThemeColor)

	ui.Render(l)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q":
			return
		case "<Down>":
			l.ScrollDown()
		case "<Up>":
			l.ScrollUp()
		case "<Enter>": // Play selected file
			selected := items[l.SelectedRow]
			AudioPlayer(vp.GetString("path")+"/"+selected, selected)
		}
		ui.Render(l)
	}
}

// Returns true if filename is a supported format
func isAudioFile(filename string) bool {
	// convert filename to lowercase to include uppercase extensions (i.e. "filename.MP3")
	filename = strings.ToLower(filename)

	for _, ext := range []string{".wav", ".mp3", ".ogg", ".flac"} {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}
