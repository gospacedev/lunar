/*
Copyright © 2022 Grantley Cullar grantcullar@gmail.com
*/
package cmd

import (
	"strings"

	in "github.com/gospacedev/lunar/cmd/internal"

	"github.com/spf13/cobra"
)

var argsString string

var newpathCmd = &cobra.Command{
	Use:   "newpath",
	Short: "Create newpath filepath",
	Long:  `Configure Lunar to the filepath that contains the mp3 files`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argsString := strings.Join(args, "")

		in.NewFilePath(argsString)
		in.Start()
	},
}

func init() {
	rootCmd.AddCommand(newpathCmd)
}
