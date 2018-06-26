package utils

import (
	"os"
	"compress/gzip"
	"archive/tar"
	log "github.com/sirupsen/logrus"
	"io"
	"strings"
	"fmt"
)

//Compress is used to compress files into tar.gz
func Compress(files []*os.File, dest string) error {
	//create the directory according to dest
	directory, _ := os.Create(dest)
	defer directory.Close()
	gzipwriter := gzip.NewWriter(directory)
	defer gzipwriter.Close()
	tarwriter := tar.NewWriter(gzipwriter)
	defer tarwriter.Close()
	for len, file := range files {
		err := compress(file, "", tarwriter)
		if err != nil {
			log.Fatalf("Failed to compress file %s: %v", file, err)
		}
		fmt.Println(len)
	}
	return nil
}

func compress(file *os.File, prefix string, tarwriter *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		log.Fatalf("%s not found: %v", file, err)
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			log.Fatal("Readdir failed: %v", err)
		}
		for _, fileInfo := range fileInfos {
			filename := file.Name() + "/" + fileInfo.Name()
			f, err := os.Open(filename)
			if err != nil {
				log.Fatalf("File to open file %s: %v", filename, err)
			}
			err = compress(f, prefix, tarwriter)
			if err != nil {
				log.Fatalf("compress failed;: %v", err)
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			log.Fatalf("tar.FileInfoHeader failed: %v", err)
		}
		header.Name = prefix + "/" + header.Name
		err = tarwriter.WriteHeader(header)
		if err != nil {
			log.Fatalf("tarwriter.WriteHeader failed: %v", err)
		}
		_, err = io.Copy(tarwriter, file)
		file.Close()
		if err != nil {
			log.Fatalf("io.Copy failed: %v", err)
		}
	}
	return nil
}

//DeCompress is used to extra files from tar.gz
func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		log.Fatalf("tar.gz file %s open failed: %v", tarFile, err)
	}
	defer srcFile.Close()
	gzipreader, err := gzip.NewReader(srcFile)
	if err != nil {
		log.Fatalf("gzip.NewReader failed: %v", err)
	}
	defer gzipreader.Close()
	tarreader := tar.NewReader(gzipreader)
	for {
		hdr, err := tarreader.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("tarreader.Next faild: %v", err)
			}
		}
		filename := dest+hdr.Name
		file, err := createFile(filename)
		if err != nil {
			log.Fatalf("createFile failed: %v", err)
		}
		io.Copy(file, tarreader)
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name,"/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
