package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func NewConfig() {
	vp := viper.New()

	// Creating json and storing path containing songs
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	// Reading config file
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	if vp.Get("path") == "" {
		fmt.Println("Enter new filepath:")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
	
		newPath := scanner.Text()
		fmt.Println(newPath)
	
		fmt.Println("The new path is " + newPath)
	
		vp.Set("path", newPath)
		vp.WriteConfig()
	}
}