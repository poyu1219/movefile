package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//srcPath := "/Volumes/Share/Download/"
	srcPath := "/Volumes/home/"

	destPath := "/Volumes/home/aa/"

	//openDir(srcPath, destPath)

	//return

	dir, err := os.Open(srcPath)
	if err != nil {
		return
	}
	defer func() {
		dir.Close()
	}()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		if fi.IsDir() {
			newDir := srcPath + fi.Name() + "/"
			if dir, err := os.Open(newDir); err == nil {
				if fileInfos, err := dir.Readdir(-1); err == nil {
					for _, fi := range fileInfos {
						if !fi.IsDir() && fi.Size() > 100000000 {
							if err := os.Rename(newDir+fi.Name(), destPath+fi.Name()); err != nil {
								fmt.Println(err)
							}
							fmt.Println(newDir + fi.Name())
							//fmt.Println(fi.Size())
						}
					}
				}
				dir.Close()
			}
		}

		//fmt.Println(fi.Name())
		//fmt.Println(destPath)
	}

}

func openDir(dirPath string, destPath string) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return
	}
	defer func() {
		dir.Close()
	}()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		if fi.IsDir() {
			newDir := dirPath + fi.Name() + "/"
			openDir(newDir, destPath)
		} else {
			if fi.Size() > 100000000 {
				filename := strings.Replace(fi.Name(), "hhd800.com@", "", 1)

				if err := os.Rename(dirPath+fi.Name(), destPath+filename); err != nil {
					fmt.Println(err)
				}
				fmt.Println(dirPath + fi.Name())
			}
		}

		//fmt.Println(fi.Name())
		//fmt.Println(destPath)
	}
}
