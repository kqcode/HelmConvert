package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
	"fmt"
	"os"
)
//Logrus hooks

//Hook for erroring and exit out on warning

type errorOnWarningHook struct {}

func (errorOnWarningHook) Levels() []log.Level {
	return []log.Level{log.WarnLevel}
}

func (errorOnWarningHook) Fire(entry *log.Entry) error {
	log.Fatalf(entry.Message)
	return nil
}

var (
	GlobalProvider string //Specifies the cloud provider
	GlobalVerbose bool //shoe info in detail
	GlobalSuppressWarnings bool //
	GlobalErrorOnWarning bool //
	GlobalFiles []string
)

var RootCmd = &cobra.Command{
	Use: "HelmCovert",
	Short: "A tool helping helm chart users to convert chart",
	Long: "HelmConvert is a tool to help helm users who are familiar with one kind of specified cloud to convert to another cloud",
	//PersistentPreRun will be "inherited" by all children and ran before *every* command unless
	//the child has overriden the functionality. This functionality was implemented to check/modify
	//all global flag calls regardless of app call.
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		//Add extra logging when verbosity is passed
		if GlobalVerbose {
			log.SetLevel(log.DebugLevel)
		}

		//Disable the timestamp (HelmConvert is too fast)
		formatter := new(log.TextFormatter)
		formatter.DisableTimestamp = true
		formatter.ForceColors = true
		log.SetFormatter(formatter)

		//set the appropriate suppress warning and error on warning flgs
		if GlobalSuppressWarnings {
			log.SetLevel(log.ErrorLevel)
		} else if GlobalErrorOnWarning{
			hook := errorOnWarningHook{}
			log.AddHook(hook)
		}

		//Error out of the user has not chosen Huawei or Ali
		provider := strings.ToLower(GlobalProvider)
		if provider != "Huawei" && provider != "Ali" {
			log.Fatalf("%s is an unsupported provider. Supported providers are: 'Huawei', 'Ali'.",GlobalProvider)
		}
	},
}

//Execute
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&GlobalVerbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().BoolVar(&GlobalSuppressWarnings, "suppress-warnings", false, "Suppress all warnings")
	RootCmd.PersistentFlags().BoolVar(&GlobalErrorOnWarning, "error-on-warning", false, "Treat any warning as an error")
	RootCmd.PersistentFlags().StringArrayVarP(&GlobalFiles, "file", "f", []string{}, "Specify the chart file to convert")
	RootCmd.PersistentFlags().StringVar(&GlobalProvider, "provider", "Huawei", "Specify a cloud provider, Huawei or Ali")
}