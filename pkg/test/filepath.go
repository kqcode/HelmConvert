package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	file := "/root/chart/wordpress-0.6.13.tgz"
	Abspath, err := filepath.Abs(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Abspath)

	dir := filepath.Dir(file)
	fmt.Println(dir)
}
