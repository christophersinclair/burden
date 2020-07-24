package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"log"
	"net/http"
  "strconv"
  "time"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

func loadtest(users int, endpoint string) string {

	for i := 0; i < users; i++ {
		go httpCall(endpoint)
	}
	return "Success"
}

func httpCall(endpoint string) {
	response, err := http.Get(endpoint)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        respCode := strconv.Itoa(response.StatusCode)
        fmt.Println("Status code for endpoint " + endpoint + ": " + respCode)
	}
}

func main() {
  greeting := `
      ..-::::::-..          
  .:-::::::::::::::-:.
  ._:::    ::    :::_.
   .:( ^   :: ^   ):.             __                   __                
   .:::   (..)   :::.            / /_  __  ___________/ /__  ____          
   .:::::::UU:::::::.           / __ \/ / / / ___/ __  / _ \/ __ \
   .::::::::::::::::.          / /_/ / /_/ / /  / /_/ /  __/ / / /
   -::::::::::::::::-         /_.___/\__,_/_/   \__,_/\___/_/ /_/ 
   .::::::::::::::::.
    .::::::::::::::.
      oO:::::::Oo`

  fmt.Println(greeting)
 
  load_users, err := strconv.Atoi(goDotEnvVariable("load_users"))
  if err != nil {
	  fmt.Printf("Failed to load user config with error %s\n", err)
  }
  http_endpoint := goDotEnvVariable("http_endpoint")

  response := loadtest(load_users, http_endpoint)
 
  // give time for the goroutines to kick off
  time.Sleep(3000 * time.Millisecond)

  fmt.Println(response)
  
}
