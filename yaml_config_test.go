package main

import (
	"fmt"
	"os"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

type yaml_nest struct {
	Desc     string
	Location struct {
		Path string
	}
}

func Test_simple_yaml(t *testing.T) {
	//		err_yml:= yaml.Unmarshal(buff, &config)
	var config yaml_test
	var example yaml_test
	buff := make([]byte, 20)
	example.Desc = "foo"

	buff, err_yaml := yaml.Marshal(example)
	if err_yaml != nil {
		fmt.Printf("err_yaml = %+v\n", err_yaml)
		t.Fatal(err_yaml)
	}

	err := yaml.Unmarshal(buff, &config)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		t.Fatal(err)
	}
	if config.Desc != "foo" {
		fmt.Printf("buff = %+v\n", buff)
		fmt.Printf("config = %+v\n", config)
		t.Fatalf("test config and config dont match")
	}
	// yaml.Unmarshal([]byte("a: 1\nb: 2"), &t)
}

func Test_nested(t *testing.T) {
	var config yaml_nest
	var example yaml_nest
	buff := make([]byte, 20)
	example.Desc = "foo"
	example.Location.Path = "path"

	buff, err_yaml := yaml.Marshal(example)
	if err_yaml != nil {
		fmt.Printf("err_yaml = %+v\n", err_yaml)
		t.Fatal(err_yaml)
	}

	err := yaml.Unmarshal(buff, &config)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		t.Fatal(err)
	}
	if config.Location.Path != "path" {
		fmt.Printf("buff = %+v\n", buff)
		fmt.Printf("config = %+v\n", config)
		t.Fatalf("test config and config dont match")
	}
}

func Test_config(t *testing.T) {
	var config Config
	var example Config
	example.Desc = "sname"
	example.Server.Name = "felix"
	example.Server.Dir = "/path"

	buff, err_yaml := yaml.Marshal(example)
	if err_yaml != nil {
		fmt.Printf("err_yaml = %+v\n", err_yaml)
	}
	if len(buff) != 47 {
		fmt.Printf("buff = %+v\n", buff)
		fmt.Println("what is the buffer like")
		fmt.Println(string(buff[:]))
		fmt.Println("length", len(buff))
		t.Fatalf("buffer with config is not complete")
	}

	err := yaml.Unmarshal(buff, &config)
	if err != nil {
		t.Fatal(err)
		fmt.Printf("err = %+v\n", err)
		fmt.Printf("config = %+v\n", config)
	}

}

func Test_read_config(t *testing.T) {
	filename := "hack.yml"
	fileinfo, fileinfoerr := os.Stat(filename)
	if fileinfoerr != nil {
		t.Fatal(fileinfoerr)
	}
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("couldnt open config file")
	}
	buff := make([]byte, fileinfo.Size())

	var config Config
	var n int
	n, err = file.Read(buff)
	if err != nil || n == 0 {
		fmt.Println("error reading the file")
	}
	file.Close()
	yaml_err := yaml.Unmarshal(buff, &config)
	if yaml_err != nil {
		fmt.Println(string(buff[:]))
		fmt.Printf("buff = %+v\n", buff)
		t.Fatal(yaml_err)
	}

	if config.Desc != "sname" {
		t.Fatalf("Tried to load config got " + config.Desc)
	}

}
