package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"log"
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

func main() {
  // godotenv package
  dotenv := goDotEnvVariable("stress_time")

  fmt.Printf("godotenv : %s = %s \n", "stress_time", dotenv)
}
