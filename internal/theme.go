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

	ui "github.com/gizak/termui/v3"
	"github.com/spf13/viper"
)

var (
	TitleTheme ui.Color
	BorderTheme ui.Color
	SelectedRowTheme ui.Color
)

func ChangeTheme(theme string) {
	switch theme {
	case "powershell":
		TitleTheme = 3
		BorderTheme = 6
		SelectedRowTheme = 3
	case "christmas":
		TitleTheme = 1
		BorderTheme = 2
		SelectedRowTheme = 3
	case "neon":
		TitleTheme = 5
		BorderTheme = 6
		SelectedRowTheme = 5
	case "cloud":
		TitleTheme = 7
		BorderTheme = 7
		SelectedRowTheme = 6
	}

	vp := viper.New()

	// get user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(home)

	// Reading config file
	err1 := vp.ReadInConfig()
	if err1 != nil {
		fmt.Println("Error: Cannot read config file")
	}

	vp.Set("TitleThemeColor", TitleTheme)
	vp.Set("BorderThemeColor", BorderTheme)
	vp.Set("SelectedRowThemeColor", SelectedRowTheme)

	vp.WriteConfig()

	fmt.Println("Theme successfully changed")
}