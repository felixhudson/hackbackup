package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type HackFile struct {
	name     string
	path     string
	modified time.Time
	size     int
	hash     string
}
type ConfigLoc struct {
	Dir string
}

type Config struct {
	Desc   string
	Server struct {
		Name string
		Dir  string
	}
	Locations []ConfigLoc
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

func loadyml(filename string) (Config, error) {
	var config Config

	dir_err := os.Chdir("C:\\Users\\Felix\\programing\\go\\src\\github.com\\user\\hackbackup\\")
	if dir_err != nil {
		fmt.Println(dir_err)
	}
	fileinfo, err := os.Stat(filename)

	if err != nil {
		fmt.Println("error couldnt get file info", filename)
		return config, err
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error couldnt open file", filename)
		return config, err
	}

	buff := make([]byte, fileinfo.Size())

	var n int
	n, err = file.Read(buff)
	if err != nil {
		fmt.Println("error reading the file")
	}
	file.Close()

	if n > 1 {
		err := yaml.Unmarshal(buff, &config)
		if err != nil {
			fmt.Println("couldnt do the yml", err)
		}
	}
	//fmt.Print(config)
	return config, nil
}


// Return servername and the directory
func get_config(filename string) (string, string) {
	var n int
	dir_err := os.Chdir("C:\\Users\\Felix\\programing\\go\\src\\github.com\\user\\hackbackup\\")
	if dir_err != nil {
		fileinfo, fileinfoerr := os.Stat(filename)

		if fileinfoerr != nil {
			fmt.Println("error couldnt get file info", filename)
			return "", ""

		}
		file, err := os.Open("hack.yml")
		fmt.Println(err)
		buff := make([]byte, fileinfo.Size())
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

func get_files(dir string) ([]HackFile , error) {
	result := make([]HackFile, 0)
	var hfile HackFile

	// look at a dir and get data
	dirdata, err := os.Open("C:\\Users\\felix")
	if err != nil {
		fmt.Println(err)
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
	//fmt.Println(os.ModeDir)
	for i := 0; i < len(ls); i++ {
		if ls[i].Mode().IsDir() {
			//fmt.Println(ls[i].Name(), "%s is a directory ")
		}
	}

	hfile.name = "filename"
	hfile.modified = time.Now()
	result = append(result, hfile)
	return result, nil
}
func get_recent_backup() ([]HackFile, error) {
	return make([]HackFile, 1), nil

}

func main() {
	// open a yml file!

	//server, dir := get_config("hack.yml")
	//fmt.Println(server, dir)

	// look at that dir and list file names and dates
	//printbytes(buff, n)
	config, err := loadyml("hack.yml")
	if err != nil {
		log.Println("Couldnt load config")
		log.Println("err")
		panic("Couldnt load config")
	}
	log.Printf("config = %+v\n", config)
	file_list, err := get_files(config.Server.Dir)
	if err != nil {
		panic("counldnt get file list")
	}
	// open current file list
	backup_set, err := get_recent_backup()
	if err != nil {
		panic("couldnt get recent backup")
	}
	compare := testable_make_list(backup_set)
	compare_list := testable_make_list(file_list)
	result, err := compare_string_file_elements(compare, compare_list)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result = %+v\n", result)

	// test connect to server

	// hash paths and names and sizes

	// (encrypt) send files

	// save backupset to disk

}

func server() {
	// initalise

	// open sockets

	// wait

	// receieve backupset
	// detect full disk
	// loop recieve files
	// calculate and update backupsets

}
