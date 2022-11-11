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

// Returns true if filename is a supported format
func isAudioFile(filename string) bool {
	filename = strings.ToLower(filename)

	for _, ext := range []string{".wav", ".mp3", ".ogg", ".flac"} {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

// Increase the length of the list to fit all items until 10 items
func DynamicHeight(items []string) int {
	if len(items) >= 10 {
		return 12
	}
	return len(items)+2
}

// Menu list all the audio files from the path set from newpath then play the selected file
func Menu() {
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

	path, err := os.ReadDir(vp.GetString("path"))
	if err != nil {
		fmt.Println("No filepath detected, to add new filepath run: lunar newpath")
		os.Exit(1)
	}

	TitleThemeColor := vp.GetInt("titlethemecolor")
	BorderThemeColor := vp.GetInt("borderthemecolor")
	SelectedRowThemeColor := vp.GetInt("SelectedRowThemeColor")

	if TitleThemeColor == 0 || BorderThemeColor == 0 {
		TitleThemeColor = 3
		BorderThemeColor = 6
		SelectedRowThemeColor = 3
	}

	var items []string

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
	l.SetRect(0, 0, 40, DynamicHeight(items))
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
		case "<Enter>":
			selected := items[l.SelectedRow]
			AudioPlayer(vp.GetString("path")+"/"+selected, selected)
		}
		ui.Render(l)
	}
}
