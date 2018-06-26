package app

import (
	"github.com/kqcode/helmconvert/pkg/convertopt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"github.com/kqcode/helmconvert/pkg/utils"
	"path/filepath"
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
	for _, chartFile := range opt.TgzFile {
		var chartDir string
		//workingDir is the directory charFile belongs to
		workingDir := filepath.Dir(chartFile)
		utils.DeCompress(chartFile, workingDir)
		//if chartFile is in absolute path format
		if filepath.IsAbs(chartFile) {
			//chartDir is the directory produces by utils.DeCompress function,
			//which is a chart directory
			chartDir = string([]rune(chartFile)[0:strings.LastIndex(chartFile,"-")])
		} else {
			//in this case, chartFile is in relative path format
			chartDir = workingDir + string([]rune(chartFile)[0:strings.LastIndex(chartFile,"-")])
		}
		//TODO how to convert
		if opt.Provider == "Huawei" {

		}
	}

}