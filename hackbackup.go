package main

import (
	"fmt"
	// "gopkg.in/yaml"
	//"io"
	"os"
)

func printbytes(data []byte, length int) {
	for i := 0; i <= length; i++ {
		//convert and print byte to ascii
		fmt.Print(fmt.Sprintf("%c", data[i]))
	}
}

// Return servername and the directory
func config(filename string) (string, string) {
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

	return "1.2.3.4", "/home/user"

}

func main() {
	// open a yml file!
	// test line zzxxyy

	file, err := os.Open("hack.yml")
	fmt.Println(err)
	buff := make([]byte, 128)
	file.Close()

	// open a normal file
	file, err = os.Open("hack.yml")
	fmt.Println(err)
	var n int
	n, err = file.Read(buff)
	fmt.Println(n)
	//fmt.Println(buff)
	printbytes(buff, n)
	fmt.Println("vim-go")
}
