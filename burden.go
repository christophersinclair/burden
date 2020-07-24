package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
  "strconv"
  "time"
  "flag"
)

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




func isFlagPassed(name string) bool {
  found := false
  flag.Visit(func(f *flag.Flag) {
    if f.Name == name {
      found = true
    }
  })

  return found
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
      oO:::::::Oo
      
  Welcome to burden! Author: @srechris
  
  You can get help by running 'burden --help'.
      `

  fmt.Println(greeting)
 

  // Declare program flags
  loadUsersPtr := flag.Int("loadusers", 10, "Number of users to simulate in a load test.")
  loadTimePtr := flag.String("loadtime", "1m", "Duration of load test. Possible times include 5m, 10s, 1hr, etc.")
  httpEndpointPtr := flag.String("httpendpoint", "", "HTTP endpoint to test. No default value.")


  // Parse flags
  flag.Parse()


  if isFlagPassed("httpendpoint") == false {
    fmt.Println("--httpendpoint not set!")
    fmt.Fprintf(os.Stderr, "error: Please provide an endpoint to test using the --httpendpoint flag.")
    os.Exit(1)
  } else {
    fmt.Println("--httpendpoint set to " + *httpEndpointPtr)
  }

  if isFlagPassed("loadusers") == false {
    fmt.Println("--loadusers not set, using default value of " + strconv.Itoa(*loadUsersPtr))
  } else {
    fmt.Println("--loadusers set to " + strconv.Itoa(*loadUsersPtr))
  }

  if isFlagPassed("loadtime") == false {
    fmt.Println("--loadtime not set, using default value of " + *loadTimePtr)
  } else {
    fmt.Println("--loadtime set to " + *loadTimePtr)
  }

  // give time for the goroutines to kick off
  time.Sleep(3000 * time.Millisecond)
  
}
