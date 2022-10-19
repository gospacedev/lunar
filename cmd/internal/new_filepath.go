package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

// Create new json file that stores filepath
func NewFilePath(newPath string) {
	vp := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	// Reading config file
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	vp.Set("path", newPath)
	vp.WriteConfig()

	fmt.Println("Filepath successfully added")
}
