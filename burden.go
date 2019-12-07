package main

import (
	"fmt"
	"flag"
)

func main() {
	confPtr := flag.String("f", "config/order.yml", "Location of order.yml")
	flag.Parse()

	config := loadconf(*confPtr)
	
	fmt.Printf("--- m:\n%v\n\n", config)

}