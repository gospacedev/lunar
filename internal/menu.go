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

	filepath, err := os.ReadDir(vp.GetString("path"))
	if err != nil {
		fmt.Println("No filepath detected, to add new filepath run: lunar newpath")
		os.Exit(1)
	}

	var files []string

	// Get file names in folder
	for _, f := range filepath {
		// Only include supported audio formats/extensions (currently only .mp3).
		// Note that this has a side-effect of also filtering out directories, which need to be
		// handled separately and navigated to instead of.. well.. crashing the application  =P
		if isAudioFile(f.Name()) {
			files = append(files, f.Name())
		}
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "Music List"
	l.Rows = files
	l.SelectedRowStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 40, len(files)+2)
	l.BorderStyle.Fg = ui.ColorCyan
	l.TitleStyle.Fg = ui.ColorYellow

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Enter>":
			selected := files[l.SelectedRow]

			// Play selected music
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
