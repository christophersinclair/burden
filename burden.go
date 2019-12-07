package main

import (
	"fmt"
	//"time"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LoadConfig struct {
	users int
}

var loadConfig LoadConfig

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadconf(filename string) []byte {
	file, err := ioutil.ReadFile(filename) // TODO: change this to a cmd-line argument
	check(err)
	return file
}

func main() {
	fmt.Println("Hello, World!")

	config := loadconf("config\\simple-order.yml")
	fmt.Println(string(config))

	yaml.Unmarshal(config, &loadConfig)
	
	fmt.Println(loadConfig)
}