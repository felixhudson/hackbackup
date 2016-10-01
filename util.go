package main

import (
		"log"
		"os"
		"io/ioutil"
		"strings"
		"fmt"

	)

type Hackfile struct {
	path string
	modified string
	hash string
}

func (h *Hackfile) generate_hash() string{
	return "123" + h.path
}

func discover_files(dir string) ([]Hackfile) {
	result := make([]Hackfile, 2)
	result = append(result, Hackfile{"p","2016/1/1","a2b6"})
	result = append(result, Hackfile{"p","2016/2/2","4fb3"})
	return result
}

func make_tree(files []Hackfile) string {
  return files[0].path
}

func compare_file_elements(base []byte, compare []byte) []string{
	// we assume both files are sorted
	// compare the first element of each file.
	// 
	lines1 := strings.Split(string(base), "\n")
	lines2 := strings.Split(string(compare), "\n")
	d1 := strings.Split(lines1[0], " ")
	d2 := strings.Split(lines2[0], " ")
	//current := 0
	//fmt.Println(base)
	//fmt.Println(lines1)
	//fmt.Println(lines1[0])
	//fmt.Println(d1)
	//fmt.Println(d1[0])
	//fmt.Println(d1[1])
	//fmt.Println("---")

	// grab first value in each
	// if the compare value is lower, it means that its a new value
	// if they are the same, then increment both pointers
	// always increment the lower one

	if d1[0] != d2[0] {
		return []string{string(d2[0])}
	}
	//fmt.Println("2---")
	d1 = strings.Split(lines1[1], " ")
	d2 = strings.Split(lines2[1], " ")
	//fmt.Println(d1, d1[0], "::")
	//fmt.Println(d2, d2[0], "::")

	if d2[0] == "" {
		return []string{""}
	}
	if d1[0] == "" {
		return []string{d2[0]}
	}

	// the compare list is different
	if d1[0] != d2[0] {
		if d1[0] > d2[0] {
			return []string{d2[0]}
		}else{
			return []string{d1[0]}
		}
	}

	fmt.Println("fell through to bottom")
	result := []string {"one", "two"}
  //return []string("one", "two")
	return result
}

func run_compare(file1 string, file2 string) {
	f1 , err:= ioutil.ReadFile(file1)
	if err != nil {
		path, _ := os.Getwd()
		log.Fatal("cant open the file", path, file1 )
	}

	f2 , err:= ioutil.ReadFile(file2)
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
	run_compare("hash_path1","hash_path2")	
}
