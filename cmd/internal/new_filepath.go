package internal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Create new json file that stores filepath
func NewFilePath() {
	vp := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	if vp.GetString("path") == "" {
		NewFilePath()
	}

	// Reading config file
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println("Error: filepath has no mp3 files")
	}


	fmt.Println("Enter new filepath...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	
	newPath := scanner.Text()
	
	vp.Set("path", newPath)
	vp.WriteConfig()

	fmt.Println("Filepath successfully added")
}