package app

import (
	"github.com/kqcode/helmconvert/pkg/convertopt"
	log "github.com/sirupsen/logrus"
	"os"
)

//ValidateChartFile validates the chart file provided for convertion
func ValidateChartFile(opt *convertopt.ConvertOptions) {
	if len(opt.InputFiles) ==0 {
		log.Fatalf("You must specify at least one chart file!")
	} else {
		for _, file := range opt.InputFiles {
			_, err := os.Stat(file)
			if err != nil {
				log.Fatalf("%s not found: %v", file, err)
			}
		}
	}
}

//Convert transforms Huawei charts to Ali charts,
//or transforms Ali charts to Huawei charts
func Convert(opt convertopt.ConvertOptions) {
	if opt.Provider == "Huawei" {

	}
}