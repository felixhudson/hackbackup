package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	//"io"
	"os"
	"time"
	"crypto/md5"
)

type HackFile struct {
	name     string
	modified time.Time
}

type Config struct {
	Desc   string
	Server struct {
		Name string
		Dir  string
	}
}

type TestConfig struct {
	Server string
	Name string
	Dir  string
}

func testmd5() {
	buff := []byte("abcdefg")
	//printhex(checksum)
	fmt.Printf("%x", md5.Sum(buff))

}

func printbytes(data []byte, length int) {
	for i := 0; i <= length; i++ {
		//convert and print byte to ascii
		fmt.Print(fmt.Sprintf("%c", data[i]))
	}
}

func loadyml(filename string) (string, string) {

	dir_err := os.Chdir("C:\\Users\\Felix\\programing\\go\\src\\github.com\\user\\hackbackup\\")
	if dir_err != nil {
		fmt.Println(dir_err)
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error couldnt open file", filename)
		return "", ""
	}

	buff := make([]byte, 64)

	var config TestConfig
	var n int
	n, err = file.Read(buff)
	if err != nil {
		fmt.Println("error reading the file")
	}
	file.Close()

	if n > 1 {
		err_yml:= yaml.Unmarshal(buff, &config)
		if err_yml != nil {
			fmt.Println("couldnt do the yml" , err_yml)
		}
	}
	fmt.Print(config)
	return "server.name.tld", "C:\\Users\\Felix"

}

func test_yaml() {
	var config Config
	var data = `
desc: test
server: t2
name: t3
dir: path/to/file 
`
	yaml.Unmarshal([]byte(data), &config)
	fmt.Printf("dump of config %v\n", config)
	printbytes([]byte(data), 49)
	// create an object and save it
	config.Desc = "example"
	config.Server.Name = "exname"
	config.Server.Dir = "edir"
	d, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Printf("%v", err)

	}
	fmt.Println(string(d))

}

// Return servername and the directory
func get_config(filename string) (string, string) {
	var n int
	dir_err := os.Chdir("C:\\Users\\Felix\\programing\\go\\src\\github.com\\user\\hackbackup\\")
	if dir_err != nil {
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
	}
	return "1.2.3.4", "C:\\User\\Felix"
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
	//printbytes(buff, n)
	server, dir = loadyml("hack.yml")
	file_list := get_files(dir)
	fmt.Println(file_list)
	testmd5()
}
