package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func (h *HackFile) generate_hash() string {
	return "123" + h.path
}

type ByPath []HackFile

func (a ByPath) Len() int           { return len(a) }
func (a ByPath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPath) Less(i, j int) bool { return a[i].path < a[j].path }

func (h *HackFile) to_s() string {
	data := fmt.Sprintf("%s/%s %s %d %s", h.path, h.name, h.modified, h.size, h.hash)
	return data
}

// Given a directory, walk it recursively and return a list of Hackfiles
func discover_files(dir string) []HackFile {
	result := make([]HackFile, 2)
	t := time.Now()
	result = append(result, HackFile{"foo", "p", t, 1000, "a2b6"})
	result = append(result, HackFile{"foo1", "p", t, 1000, "4fb3"})
	return result
}

func make_tree(files []HackFile) string {
	return files[0].path
}

func compare_string_file_elements(base []string, compare []string) ([]string, error) {
	newbase := make([]byte, 0)
	for _, element := range base {
		newbase = append(newbase, []byte("\n")...)
		newbase = append(newbase, []byte(element)...)
	}
	newcompare := make([]byte, 0)
	for _, element := range base {
		newcompare = append(newcompare, []byte("\n")...)
		newcompare = append(newcompare, []byte(element)...)
	}
	result, err := compare_file_elements(newbase, newcompare)
	if err != nil {
		return make([]string, 0), err
	}
	return result, nil
}

func compare_file_elements(base []byte, compare []byte) ([]string, error) {
	// we assume both files are sorted
	// compare the first element of each file.
	//
	var lines1, lines2 []string
	var d1, d2 []string
	end := len(lines2) + 1
	current := 0
	counter := 0

	if !(bytes.Contains(base, []byte("\n")) || bytes.Contains(compare, []byte("\n"))) {
		return make([]string, 0), errors.New("Data is malformed")
	}

	lines1 = strings.Split(string(base), "\n")
	lines2 = strings.Split(string(compare), "\n")
	result := make([]string, 0)
	for current <= end {
		d1 = strings.Split(lines1[counter], " ")
		d2 = strings.Split(lines2[current], " ")
		// grab first value in each
		// if the compare value is lower, it means that its a new value
		// if they are the same, then increment both pointers
		// always increment the lower one
		//log.Printf("d1 = %+v\n", d1)
		//log.Printf("d2 = %+v\n", d2)
		//log.Println("---")

		// deal with empty paths
		if d2[0] == "/" {
			current++
			continue
		}

		if d1[0] == "/" {
			counter++
			continue
		}

		if d2[0] == "" {
			//result = append(result, "")
		}
		if d1[0] == "" {
			result = append(result, d2[0])
			current++
			continue
		}

		// the compare list is different
		if d1[0] != d2[0] {
			if d1[0] > d2[0] {
				result = append(result, d2[0])
				current++
				//return []string{d2[0]}
			} else {
				//return []string{d1[0]}
			}
		} else {
			counter++
			current++
		}
	}
	return result, nil
}

func run_compare(file1 string, file2 string) {
	f1, err := ioutil.ReadFile(file1)
	if err != nil {
		path, _ := os.Getwd()
		log.Fatal("cant open the file", path, file1)
	}

	f2, err := ioutil.ReadFile(file2)
	if err != nil {
		log.Fatal("cant open the file")
	}
	//buff := make([]byte,50)
	//n, readerr := f1.Read(buff)

	compare_file_elements(f1, f2)
}

func make_backup(dir string) {
	// put together a tree of nodes representing the backup

	// find all files
	files := discover_files("testdir")
	// construct tree
	make_tree(files)
	// compare this to last backup
	// send tree to location
	// send new files to location

}

func make_backup_lists(dir string) {

	// alternative plan
	// make a list of all files
	// store them in a file sorted by path
	// also store them sorted by hash
	// thus we can tell which files need to be backed up easily
	// we can also work out which files are not needed.
	// simply walk through each list one at a time, finding missing entries as detected

	// lets test walking the two files
	run_compare("hash_path1", "hash_path2")
	files := discover_files(dir)
	// save the file as a list of things
	sort.Sort(ByPath(files))
	destination := "/tmp/backupset"
	fd, err := os.Open(destination)
	if err != nil {
		log.Println("Couldnt open file to write backupset")
		log.Printf("err = %+v\n", err)
	}

	for _, element := range files {
		fd.WriteString(element.to_s())
	}
	fd.Close()

}

func testable_make_list(files []HackFile) []string {
	//files := discover_files(dir)
	// save the file as a list of things
	sort.Sort(ByPath(files))
	result := make([]string, 0)
	for _, element := range files {
		result = append(result, element.to_s())
	}
	return result
}


func make_filename() string {
	date := time.Now()
	sdate := "hackbackup-" + date.Format("2006-01-02-150405")
	return sdate
}

func save_backupset_disk(backupset []HackFile) (string, error) {
	filedata := testable_make_list(backupset)

	path := "/Users/Felix/" + make_filename()
	//path = "/Users/Felix/test"
	file, err := os.Create(path)
	if err != nil {
		log.Printf("path = %+v\n", path)
		log.Println("Couldnt Open file for storage")
		log.Printf("err = %+v\n", err)
		return "", err
	}

	for _, element := range(filedata) {
		n, err := file.WriteString(element)
		if err != nil || n == 0{
			log.Println("Couldnt write backupset to disk")
			fmt.Println(err)
			return "", err
		}
	}

	return "path", nil
}
