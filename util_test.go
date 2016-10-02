package main
import "testing"
import "fmt"
import "time"

func Test_compare_file_elements(t *testing.T) {
	a := []byte("1234 bar\n1345 foo\n")
	b := []byte("1234 bar \n")
	data := compare_file_elements(a,b)
	expected := ""
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s\n)" , expected, data[0])
		t.Error()

	}
}
func Test_compare_file_elements2(t *testing.T) {
	a := []byte("1234 bar \n")
	b := []byte("1234 bar \n1345 foo\n")
	data := compare_file_elements(a,b)
	expected := "1345"
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s ::\n)" , expected, data[0])
		t.Error()
	}
}

func Test_compare_file_elements3(t *testing.T) {
	a := []byte("1234 bar \n2222 nope")
	b := []byte("1234 bar \n1345 foo\n2222 nope")
	data := compare_file_elements(a,b)
	expected := "1345"
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s ::\n)" , expected, data[0])
		t.Error()

	}
}

func Test_compare_file_many(t *testing.T) {
	a := []byte("1234 bar \n2222 nope\n3333 nope")
	b := []byte("1234 bar \n1345 foo\n2222 nope\n3334 yup")
	data := compare_file_elements(a,b)
	expected := "1345"
	if data[0] != expected {
		fmt.Printf("Expected %s, got %s ::\n)" , expected, data[0])
		t.Error()

	}
}
func assert_equal(t *testing.T, expected string, data string) {
	if data != expected {
		fmt.Printf("Expected '%s', got '%s'\n)" , expected, data)
		t.Error()
	}
}

func Test_create_fileset(t *testing.T) {
	var mock_dirlist []HackFile
	mock_dirlist = discover_files("/")
	compare := string(mock_dirlist[0].name)
  assert_equal(t, "one",string(mock_dirlist[0]))
}
