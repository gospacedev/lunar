package internal

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

// The menu where the user can choose what audio file to play
func Start() {
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

	prompt := promptui.Select{
		Label: "Select music",
		Items: files,
	}

	_, selected, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	// Play selected music
	AudioPlayer(vp.GetString("path")+"/"+selected, selected)
}
