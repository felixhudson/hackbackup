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
	fmt.Printf("err_yaml = %+v\n", err_yaml)
	fmt.Printf("buff = %+v\n", buff)

	err := yaml.Unmarshal(buff, &config)
	fmt.Printf("err = %+v\n", err)
	fmt.Printf("config = %+v\n", config)
	// yaml.Unmarshal([]byte("a: 1\nb: 2"), &t)
}

func Test_nested(t *testing.T) {
	var config yaml_nest
	var example yaml_nest
	buff := make([]byte, 20)
	example.Desc = "foo"
	example.Location.Path = "path"

	buff, err_yaml := yaml.Marshal(example)
	fmt.Printf("err_yaml = %+v\n", err_yaml)
	fmt.Printf("buff = %+v\n", buff)

	err := yaml.Unmarshal(buff, &config)
	fmt.Printf("err = %+v\n", err)
	fmt.Printf("config = %+v\n", config)
}

func Test_config(t *testing.T) {
	var config Config
	var example Config
	example.Desc = "sname"
	example.Server.Name = "felix"
	example.Server.Dir = "/path"

	buff, err_yaml := yaml.Marshal(example)
	fmt.Printf("err_yaml = %+v\n", err_yaml)
	fmt.Printf("buff = %+v\n", buff)
	fmt.Println("what is the buffer like")
	fmt.Println(string(buff[:]))

	err := yaml.Unmarshal(buff, &config)
	fmt.Printf("err = %+v\n", err)
	fmt.Printf("config = %+v\n", config)

}

func Test_read_config(t *testing.T) {
	filename := "hack.yml"
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("couldnt open config file")
	}
	buff := make([]byte, 51)

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
