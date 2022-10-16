package main

import (
	"fmt"
	"log"
	"os"

	in "lunar/internal"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func main() {
	vp := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	in.NewFilePath()

	// Reading config file
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println("Error: cannot read config file")
	}

	filepath, err := os.ReadDir(vp.GetString("path"))
	if err != nil {
        log.Fatal(err)
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
	in.MusicPlayer(vp.GetString("path") + "/" + result)

}