package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func main() {
	vp := viper.New()

	// Creating json and storing path containing songs
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	/* If the path is empty ask the user for input and save it */
	//if vp.GetString("path") == "" {
		newFilePath()
	//}

	// Reading config file
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	filepath, err := os.ReadDir(vp.GetString("path"))
	if err != nil {
        log.Fatal(err)
    }

	var files []string

	// Get all the mp3 files in a directory named music in current directory
	for _, f := range filepath {
		files = append(files, f.Name())
	}

	prompt := promptui.Select{
		Label: "Select music",
		Items: files,
	}

	// result is the item slected
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Play selected music
	mp3Player(vp.GetString("path") + "/" + result)

}