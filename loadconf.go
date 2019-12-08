package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

type BurdenConfig struct {
	Load struct {
		users int
		requests int
		HTTP struct {
			Endpoint struct {
				url string
				method string
			}
		}
	}

	Stress struct {
		users int
		requests int
		HTTP struct {
			Endpoint struct {
				url string
				method string
			}
		}
	}

	Spike struct {
		users int
		requests int
		HTTP struct {
			Endpoint struct {
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

	cm := BurdenConfig{}
	err = yaml.Unmarshal([]byte(conf), &cm)
	check(err)
	fmt.Printf("--- cm:\n%v\n\n", cm)

	return m

}