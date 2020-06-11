package main

import (
	"log"

	"./rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI(":8000"))
}
