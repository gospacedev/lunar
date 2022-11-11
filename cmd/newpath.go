/*
Copyright © 2022 Grantley Cullar grantcullar@gmail.com
*/
package cmd

import (
	"strings"

	"github.com/gospacedev/lunar/internal"

	"github.com/spf13/cobra"
)

var newpathCmd = &cobra.Command{
	Use:   "newpath",
	Short: "Create new filepath",
	Long:  `Configure Lunar to the filepath that contains the mp3 files`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argsString := strings.Join(args, "")

		internal.NewFilepath(argsString)
	},
}

func init() {
	rootCmd.AddCommand(newpathCmd)
}
