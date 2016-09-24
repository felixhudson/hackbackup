package main

import (
	"fmt"
	// "gopkg.in/yaml"
	//"io"
	"os"
	"time"
)

type HackFile struct {
	name     string
	modified time.Time
}

func printbytes(data []byte, length int) {
	for i := 0; i <= length; i++ {
		//convert and print byte to ascii
		fmt.Print(fmt.Sprintf("%c", data[i]))
	}
}

// Return servername and the directory
func get_config(filename string) (string, string) {
	var n int
	file, err := os.Open("hack.yml")
	fmt.Println(err)
	buff := make([]byte, 128)
	file.Close()

	n, err = file.Read(buff)
	for i := 0; i < n; i++ {
		if buff[i] == 66 {
			fmt.Println("error")
		}
	}
	return "1.2.3.4", "C:\\User\\felix"
}

func get_files(dir string) []HackFile {
	result := make([]HackFile, 0)
	var hfile HackFile

	// look at a dir and get data
	dirdata, err := os.Open("C:\\Users\\felix")
	fmt.Println(err)
	if err != nil {
		fmt.Println("couldnt read directory")
	}
	ls, err := dirdata.Readdir(0)
	if err != nil {
		fmt.Println("couldnt read directory listing")
	}

	for i := 0; i < len(ls); i++ {
		//		fmt.Println(ls[i].Name())
		//		fmt.Println(ls[i].ModTime())
		hfile.name = ls[i].Name()
		hfile.modified = ls[i].ModTime()
		result = append(result, hfile)
	}
	// find all directories in this dir
	fmt.Println(os.ModeDir)
	for i := 0; i < len(ls); i++ {
		if ls[i].Mode().IsDir() {
			fmt.Println(ls[i].Name(), "%s is a directory ")
		}
	}

	hfile.name = "filename"
	hfile.modified = time.Now()
	result = append(result, hfile)
	return result
}

func main() {
	// open a yml file!

	server, dir := get_config("hack.yml")
	fmt.Println(server, dir)

	// look at that dir and list file names and dates
	file_list := get_files(dir)
	fmt.Println(file_list)
	//printbytes(buff, n)
	fmt.Println("vim-go")
}
