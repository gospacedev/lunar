package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func Start() {
	vp := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	// Reading config file
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println("Run lunar add")
	}

	filepath, err := os.ReadDir(vp.GetString("path"))
	if err != nil {
		log.Fatal("No file path delected: run lunar add")
    }

	var files []string

	// Get file names in folder
	for _, f := range filepath {
		files = append(files, f.Name())
	}

	prompt := promptui.Select{
		Label: "Select music:",
		Items: files,
	}

	// result is the item slected
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Play selected music
	MusicPlayer(vp.GetString("path") + "/" + result)

}