package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Create new json file that stores filepath
func NewFilePath(newPath string) {
	vp := viper.New()

	// get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	// stub init
	configHome := home
    configName := "config"
	configType := "json"
	configPath := filepath.Join(configHome, configName+"."+configType)
// ----

	vp.AddConfigPath(configHome)
	vp.SetConfigName(configName)
	vp.SetConfigType(configType)

	_, err2 := os.Stat(configPath)
	if !os.IsExist(err2) {
		if _, err2 := os.Create(configPath); err2 != nil {
			fmt.Print()// ignore error
		}
	}

	// Reading config file
	err1 := vp.ReadInConfig()
	if err1 != nil {
		fmt.Print()
	}

	vp.Set("path", newPath)
	vp.WriteConfig()

	fmt.Println("Filepath successfully added")
}
