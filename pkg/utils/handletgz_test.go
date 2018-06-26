package utils

import (
	"testing"
	//"os"
	//log "github.com/sirupsen/logrus"
)

func TestCompress(t *testing.T) {
	//var files []*os.File
	//file, err := os.Open("/root/chart/wordpress")
	//if err != nil {
	//	log.Fatalf("os.Open failed: %v", err)
	//}
	//files = append(files, file)
	//dest := "/root/chart/wordpress.tgz"
	//Compress(files, dest)

	deFile := "/root/chart/wordpress-0.6.13.tgz"
	deDest := "/root/chart/"
	DeCompress(deFile, deDest)


}
