package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"
	"testing"
	"time"
	"math/rand"
)

func Test_compare_string_file_elements(t *testing.T) {
	base := make([]string, 0)
	base = append(base, "one")
	data, err := compare_string_file_elements(base, base)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 && data[0] != "" {
		t.Fatal("data is borked")
	}
}

func Test_should_deal_with_bad_data(t *testing.T) {
	a := []byte("1234 bar no newline")
	b := []byte("1234 bar no newline")
	if !bytes.Contains([]byte("1234\n"), []byte("\n")) {
		t.Fatal("Error, compare doesnt work")
	}
	result, err := compare_file_elements(a, b)
	// only if we dont get an error do we fail the test
	if err == nil {
		t.Fatal(err)
	}
	if len(result) > 0 {
		t.Fatal("Expected no differences")
	}
}

func Test_compare_file_strings(t *testing.T) {
	one := []byte{'a',' ','b',' ','c', '\n'}
	two := []byte{'a',' ','b',' ','c', '\n'}
	difference, err := compare_file_elements(one,two)
	if err != nil {
		t.Fatal(err)
	}
	if len(difference) != 0 {
		log.Printf("difference = %+v\n", difference)
		t.Fatal("Difference expected to be 0, got ", len(difference))
	}
	one = []byte{'a',' ','b',' ','c', '\n'}
	two = []byte{'a',' ','b',' ','c', '\n'}
	difference, err = compare_file_elements(one,two)
	if err != nil {
		t.Fatal(err)
	}
	if len(difference) != 0 {
		t.Fatal("Difference expected to be 0")
	}
}
func Test_compare_file_elements_simple(t *testing.T) {
	a := []byte("1234 bar\n")
	b := []byte("1234 bar\n")
	data, err := compare_file_elements(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 0 {
		fmt.Printf("Expected nil data being returned")
		t.Fatal()
	}
}

func Test_compare_file_elements(t *testing.T) {
	// as file is missing in b it wont be returned!
	a := []byte("1234 bar\n1345 foo\n")
	b := []byte("1234 bar\n")
	data, err := compare_file_elements(a, b)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 0 {
		fmt.Printf("Expected nil data being returned")
		log.Printf("data = %+v\n", data)
		t.Fatal()
	}
}

func Test_compare_file_elements_new_file(t *testing.T) {
	// new file in seccond set of files
	a := []byte("1234 bar\n")
	b := []byte("1234 bar\n1345 foo\n")
	data, err := compare_file_elements(a, b)
	expected := "1345"
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 1 {
		log.Println("Expected 1 element being returned")
		log.Printf("Got = %+v\n", len(data))
		log.Printf("data = %+v\n", data)
		t.Fatal()
	}
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s\n)", expected, data[0])
		t.Error("Expected ", expected, "got", data[0])
	}
}
func Test_compare_file_elements2(t *testing.T) {
	a := []byte("1234 bar \n")
	b := []byte("1234 bar \n1345 foo\n")
	data, _ := compare_file_elements(a, b)
	expected := "1345"
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s ::\n)", expected, data[0])
		t.Error()
	}
}

func Test_compare_file_elements3(t *testing.T) {
	a := []byte("1234 bar \n2222 nope")
	b := []byte("1234 bar \n1345 foo\n2222 nope")
	data, err := compare_file_elements(a, b)
	if err != nil {
		t.Fatal(err)
	}
	expected := "1345"
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s ::\n)", expected, data[0])
		t.Error()
	}
}

func Test_compare_file_many(t *testing.T) {
	a := []byte("1234 bar \n2222 nope\n3333 nope")
	b := []byte("1234 bar \n1345 foo\n2222 nope\n3334 yup")
	data, err := compare_file_elements(a, b)
	if err != nil {
		t.Fatal(err)
	}
	expected := "1345"
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s ::\n)", expected, data[0])
		t.Error()
	}
}
func assert_equal(t *testing.T, expected string, data string) {
	if data != expected {
		fmt.Printf("Expected '%s', got '%s'\n)", expected, data)
		t.Error()
	}
}

func Test_create_fileset(t *testing.T) {
	var mock_dirlist []HackFile
	mock_dirlist = append(mock_dirlist, HackFile{"foo", "p", time.Now(), 1000, "ab34"})
	mock_dirlist = append(mock_dirlist, HackFile{"bar", "p", time.Now(), 1000, "ab34"})
	mock_dirlist = append(mock_dirlist, HackFile{"foo", "/q", time.Now(), 1000, "ab34"})
	compare := string(mock_dirlist[0].name)
	assert_equal(t, "foo", compare)
}

func Test_make_list(t *testing.T) {
	var mock_dirlist []HackFile
	mock_dirlist = append(mock_dirlist, HackFile{"foo", "p", time.Now(), 1000, "ab34"})
	mock_dirlist = append(mock_dirlist, HackFile{"bar", "p", time.Now(), 1000, "ab34"})
	mock_dirlist = append(mock_dirlist, HackFile{"foo", "/q", time.Now(), 1000, "ab34"})
	foo := testable_make_list(mock_dirlist)
	compare := "bar"
	if strings.Contains(foo[0], compare) {
		t.Error("Couldnt make backup list")
	}
}
func Test_make_big_list(t *testing.T) {
	files := discover_files("~/")
	backupset := testable_make_list(files)
	if backupset == nil {
		t.Fatal("Couldnt make list from home directory")
	}
}

func Test_compare_identical(t *testing.T) {

	files := discover_files("~/")
	backupset := testable_make_list(files)
	files = discover_files("~/")
	backupset2 := testable_make_list(files)
	//fmt.Printf("backupset = %+v\n", backupset)
	//fmt.Printf("backupset = %+v\n", backupset2)
	target_files, err := compare_string_file_elements(backupset, backupset2)
	if err != nil {
		t.Fatal(err)
	}

	if len(target_files) > 0 && target_files[0] != "" {
		t.Fatal("Error, sets should be the same")
	}

	// save lists to disk then compare
	var path string
	path, err = save_backupset_disk(files)
	if err != nil {
		log.Printf("path = %+v\n", path)
		t.Fatal(err)
	}
}

func Test_sort_list(t *testing.T) {
	mock_dirlist := make([]HackFile, 0)
	mock_dirlist = append(mock_dirlist, HackFile{"foo", "p", time.Now(), 1000, "ab34"})
	mock_dirlist = append(mock_dirlist, HackFile{"bar", "/a/a", time.Now(), 1000, "ab34"})
	mock_dirlist = append(mock_dirlist, HackFile{"foo", "/a/b", time.Now(), 1000, "ab34"})
	// sort the slice
	sort.Sort(ByPath(mock_dirlist))
	compare := "bar"
	if mock_dirlist[0].name != compare {
		fmt.Printf("compare = %+v\n", compare)
		fmt.Printf("mock_dirlist = %+v\n", mock_dirlist)
		t.Error("didnt sort")
	}
}

func Test_make_filename(t *testing.T) {
	data := make_filename()
	if len(data) != 28 {
		t.Log(len(data))
		t.Fatal("problem with the date generator")
	}
}

func Test_runtwobackups(t *testing.T) {
	backup := make([]HackFile, 0)
	for i := 0; i < 3; i++ {
		backup = append(backup,mock_file())
	}
	backupset := testable_make_list(backup)
	//alter := rand.Intn(len(backup))
	newfile := mock_file()
	// line bellow will alter the array
	//backup[alter] = newfile
	backup = append(backup, newfile)
	backupset2 := testable_make_list(backup)
	result, err := compare_string_file_elements(backupset,backupset2)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 1 {
		fmt.Printf("backupset = %+v\n", backupset)
		fmt.Printf("backupset2 = %+v\n", backupset2)
		fmt.Printf("result = %+v\n", result)
		t.Fatal("Expected 1 file got", len(result))
	}
}
func Test_runtwobackups_sort(t *testing.T) {
	backup := make([]HackFile, 0)
	backup = append(backup,mock_file())
	backupset := testable_make_list(backup)
	last := "File"
	for _, value := range backupset {
		if strings.Split(" ",value)[0] > last {
			log.Printf("backupset = %+v\n", backupset)
			log.Printf("value = %+v\n", value)
			log.Printf("last = %+v\n", last)
			t.Fatal("List of length 1 is out of order?")
		}
	}
	for i := 0; i < 3; i++ {
		backup = append(backup,mock_file())
	}
	backupset = testable_make_list(backup)

	last = "0"
	for _, value := range backupset {
		if strings.Split(" ",value)[0] > last {
			log.Printf("backupset = %+v\n", backupset)
			t.Fatal("sorted values are out of order")
		}
	}
}

func Test_runtwobackups_with_alter(t *testing.T) {
	backup := make([]HackFile, 0)
	for i := 0; i < 3; i++ {
		backup = append(backup,mock_file())
	}
	backupset := testable_make_list(backup)
	alter := rand.Intn(len(backup))
	newfile := mock_file()
	// line bellow will alter the array
	backup[alter] = newfile
	backupset2 := testable_make_list(backup)
	result, err := compare_string_file_elements(backupset,backupset2)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 1 {
		fmt.Printf("backupset = %+v\n", backupset)
		fmt.Printf("backupset2 = %+v\n", backupset2)
		fmt.Printf("result = %+v\n", result)
		t.Fatal("Expected 1 file got", len(result))
	}
}

func mock_file() HackFile {
	pick := rand.Intn(100)
	d := fmt.Sprintf("%d", pick)
	name := "File-" + d + "xx"
	return HackFile{name, "path", time.Now(), 1234, "hash"}
}

/*
test ideas
run one backup change files then run again
run two backups then run the clean up
make a loop and do it 100 times! backup and clean...
*/
