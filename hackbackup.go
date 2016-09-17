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

func main() {
	// open a yml file!
	// test line zzxxyy

	// open a normal file
	file, err := os.Open("hack.yml")
	fmt.Println(err)
	buff := make([]byte, 128)
	var n int
	n, err = file.Read(buff)
	fmt.Println(n)
	//fmt.Println(buff)
	printbytes(buff, n)
	fmt.Println("vim-go")
}
