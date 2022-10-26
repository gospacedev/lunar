package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// The menu where the user can choose what audio file to play
func Menu() {
	vp := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	// Reading config file
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	filepath, err := os.ReadDir(vp.GetString("path"))
	if err != nil {
		fmt.Println("No filepath detected: lunar newpath")
		os.Exit(0)
	}

	var files []string

	// Get file names in folder
	for _, f := range filepath {
		files = append(files, f.Name())
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "Music List"
	l.Rows = files
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 40, len(files)+2)
	l.BorderStyle.Fg = ui.ColorCyan

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
