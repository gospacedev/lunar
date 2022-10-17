/*
Copyright © 2022 Grantley Cullar grantcullar@gmail.com

*/
package cmd

import (
	in "lunar/cmd/internal"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new filepath",
	Long: `Configure Lunar to the filepath that contains the mp3 files`,
	Run: func(cmd *cobra.Command, args []string) {
		in.NewFilePath()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
