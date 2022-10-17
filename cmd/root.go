/*
Copyright © 2022 Grantley Cullar grantcullar@gmail.com

*/
package cmd

import (
	"os"

	in "github.com/gospacedev/lunar/cmd/internal"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lunar",
	Short: "Lunar is a CLI Music Player",
	Long:  `Lunar is a CLI Music Player for terminal enthusiasts`,
	Run: func(cmd *cobra.Command, args []string) {
		in.Start()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
