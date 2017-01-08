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
func (a ByPath) Less(i, j int) bool { 
	return a[i].path+"/"+a[i].name < a[j].path+"/"+a[j].name
}

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
		newbase = append(newbase, []byte(element)...)
		newbase = append(newbase, []byte("\n")...)
	}
	newcompare := make([]byte, 0)
	for _, element := range compare {
		newcompare = append(newcompare, []byte(element)...)
		newcompare = append(newcompare, []byte("\n")...)
	}
	result, err := compare_file_elements(newbase, newcompare)
	if err != nil {
		log.Println("We have a problem when comparing the files")
		log.Printf("err = %+v\n", err)
		return make([]string, 0), err
	}
	return result, nil
}

func compare_file_elements(base []byte, compare []byte) ([]string, error) {
	// we assume both files are sorted
	// compare the first element of each file.
	//
	var lines1, lines2 []string
	var split_line []string
	debug := false
	var exists bool
	var compare_value string
	basemap := make(map[string]string)


	if !(bytes.Contains(base, []byte("\n")) || bytes.Contains(compare, []byte("\n"))) {
		return make([]string, 0), errors.New("Data is malformed")
	}

	lines1 = strings.Split(string(base), "\n")
	lines2 = strings.Split(string(compare), "\n")
	p := "foo"
	for _, val := range(lines1) {
		split_line = strings.Split(val, " ")
		p = split_line[0]
		basemap[p] = val
	}
	if debug {
		log.Println("############")
		log.Printf("lines1 = %+v\n", strings.Join(lines1,"\n"))
		log.Printf("lines2 = %+v\n", strings.Join(lines2,"\n"))
		log.Printf("basemap = %+v\n", basemap)
	}
	result := make([]string, 0)
	for _, val := range(lines2) {
		split_line = strings.Split(val, " ")
		p = split_line[0]
		compare_value, exists = basemap[p]
		if !exists {
			if debug {
				log.Println("Adding to result", compare_value)
			}
			// it doesnt exist in the base thus we include it
			result = append(result, p)
		} else {
			if debug {
				log.Println("Comparing", p)
				log.Println("found", compare_value)
			}
		}
	}
	return result, nil
	/*
	return make([]string,0), errors.New("not a real error")
	end := len(lines1)
	end2 := len(lines2)
	if debug {
		log.Printf("lines1 = %+v\n", lines1)
		log.Printf("lines2 = %+v\n", lines2)
	}
	for (pos2 < end2 - 1){
		if debug {
			log.Println("---")
			log.Printf("counter = %+v\n", counter)
			log.Printf("pos2 = %+v\n", pos2)
			log.Printf("end = %+v\n", end)
			log.Printf("end2 = %+v\n", end2)
		}

		//if list1 is at end just take d2
		if counter == end {
			if debug {
				log.Println("we are adding d2 as list 1 is at end")
			}
			result = append(result, d2[0])
			pos2++
			continue
		}

		d1 = strings.Split(lines1[counter], " ")
		d2 = strings.Split(lines2[pos2], " ")
		// grab first value in each
		// if the compare value is lower, it means that its a new value
		// if they are the same, then increment both pointers
		// always increment the lower one
		if debug {
			log.Printf("d1 = %+v\n", d1)
			log.Printf("d2 = %+v\n", d2)
		}

		// deal with empty paths
		if d2[0] == "/" {
			pos2++
			continue
		}

		if d1[0] == "/" {
			counter++
			continue
		}

		if d2[0] == "" {
			if debug {
				log.Println("Skipping d2 as it is empty")
			}
			pos2++
			continue
		}

		if d1[0] == "" {
			if debug {
				log.Println("we are skipping d1 as it is empty")
			}
			counter++
			continue
		}

		// the compare list is different
		if d1[0] != d2[0] {
			if d1[0] > d2[0] {
				// thus d2 is new and appears in the list before the value of d1
				if debug {
					log.Println("we are taking d2 as its not in list1, moving to next value of list2")
				}
				result = append(result, d2[0])
				pos2++
				continue
			} else {
				// thus d2 is new and appeares after the value of d1
				// so we move d1 to the next value
				if debug {
					log.Println("we are skipping d1 as its different but occurs after d2")
				}
				counter ++
				continue
			}
		} else {
			if debug {
				log.Println("paths are the same", counter, pos2)
			}
			// only increment the second list
			pos2++
			continue
		}
		if debug {
			log.Println("dropped to bottom")
		}
	}
	if debug {
		log.Println("end of loop compare")
	}
	return result, nil
	*/
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
