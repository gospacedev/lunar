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

	// Reading config file
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Enter new filepath...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	newPath := scanner.Text()

	vp.Set("path", newPath)
	vp.WriteConfig()

	fmt.Println("Filepath successfully added")
}
