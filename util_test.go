package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"
)

func Test_compare_string_file_elements(t *testing.T) {
	base := make([]string, 0)
	base = append(base, "one")
	data, err := compare_string_file_elements(base, base)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) >0 && data[0] != "" {
		t.Fatal("data is borked")
	}
}

func Test_should_deal_with_bad_data(t *testing.T) {
	a := []byte("1234 bar no newline")
	b := []byte("1234 bar no newline")
	if ! bytes.Contains([]byte("1234\n"), []byte("\n")) {
		t.Fatal("Error, compare doesnt work")
	}
	_, err := compare_file_elements(a, b)
	// only if we dont get an error do we fail the test
	if err == nil {
		t.Fatal(err)
	}
}

func Test_compare_file_elements(t *testing.T) {
	a := []byte("1234 bar\n1345 foo\n")
	b := []byte("1234 bar \n")
	data, _ := compare_file_elements(a, b)
	expected := ""
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s\n)", expected, data[0])
		t.Error()

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
	data, _ := compare_file_elements(a, b)
	expected := "1345"
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s ::\n)", expected, data[0])
		t.Error()

	}
}

func Test_compare_file_many(t *testing.T) {
	a := []byte("1234 bar \n2222 nope\n3333 nope")
	b := []byte("1234 bar \n1345 foo\n2222 nope\n3334 yup")
	data, _ := compare_file_elements(a, b)
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
