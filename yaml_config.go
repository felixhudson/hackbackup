package main

//import yaml "gopkg.in/yaml.v2"

type yaml_test struct {
	Desc string
}

func load_config(path string) Config {

	var foo Config
	foo.Desc = "blah"
	//yaml.Marshal()
	return foo

}
