package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type BurdenConfig struct {
	Test struct {
		users int
		http struct {
			endpoint struct {
				url string
				method string
			}
		}
	}
}

var burdenConfig BurdenConfig

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadconf(filename string) map[interface{}]interface{} {
	conf, err := ioutil.ReadFile(filename)
	check(err)
	
	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(conf), &m)
	check(err)

	return m

}