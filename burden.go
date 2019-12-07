package main

import (
	"fmt"
	//"time"
)

func main() {
	fmt.Println("Hello, World!")

	config := loadconf("config\\simple-order.yml") // TODO: change to cmd-line argument (-f order.yml)
	
	fmt.Printf("--- m:\n%v\n\n", config)

}