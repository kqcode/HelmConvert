package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kqcode/helmconvert/pkg/app"
	"github.com/kqcode/helmconvert/pkg/convertopt"
	"github.com/spf13/viper"
)

var ConvertOpt convertopt.ConvertOptions

var convertCmd = &cobra.Command {
	Use: "convert [file]",
	Short: "Convert a specified chart for another cloud provider",
	PreRun: func(cmd *cobra.Command, args []string) {
		ConvertOpt = convertopt.ConvertOptions {
			InputFiles: GlobalFiles,
			Provider: GlobalProvider,
			TgzFile: GlobalFiles,
		}
		app.ValidateChartFile(&ConvertOpt)
	},
	Run: func(cmd *cobra.Command, args []string) {
		app.Convert(ConvertOpt)
	},
}

func init() {
	//Automatically grab environments variables
	viper.AutomaticEnv()
}
