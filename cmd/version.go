package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kqcode/helmconvert/pkg/version"
	"fmt"
)


var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version of HelmConvert",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.VERSION + " (" + version.GITCOMMIT + ")")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}