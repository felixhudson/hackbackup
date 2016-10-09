package main

import "testing"
import "fmt"
import "sort"
import "strings"
import "time"

func Test_compare_file_elements(t *testing.T) {
	a := []byte("1234 bar\n1345 foo\n")
	b := []byte("1234 bar \n")
	data := compare_file_elements(a, b)
	expected := ""
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s\n)", expected, data[0])
		t.Error()

	}
}
func Test_compare_file_elements2(t *testing.T) {
	a := []byte("1234 bar \n")
	b := []byte("1234 bar \n1345 foo\n")
	data := compare_file_elements(a, b)
	expected := "1345"
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s ::\n)", expected, data[0])
		t.Error()
	}
}

func Test_compare_file_elements3(t *testing.T) {
	a := []byte("1234 bar \n2222 nope")
	b := []byte("1234 bar \n1345 foo\n2222 nope")
	data := compare_file_elements(a, b)
	expected := "1345"
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s ::\n)", expected, data[0])
		t.Error()

	}
}

func Test_compare_file_many(t *testing.T) {
	a := []byte("1234 bar \n2222 nope\n3333 nope")
	b := []byte("1234 bar \n1345 foo\n2222 nope\n3334 yup")
	data := compare_file_elements(a, b)
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
