package main

import (
	"fmt"
	"net/http"
)

func httpCall(endpoint string) {
	response, err := http.Get(endpoint)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        fmt.Println("Status for endpoint " + endpoint + ": " + response.Status)
	}

}