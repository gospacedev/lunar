/*
Copyright © 2022 Grantley Cullar <grantcullar@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Create new json file that stores filepath
func NewFilepath(newPath string) {
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
			fmt.Print() // ignore error
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
