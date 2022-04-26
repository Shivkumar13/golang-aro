package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	fmt.Println("Hello, world!")

	f, err := os.OpenFile("golang-demo.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer f.Close()

	log.SetOutput(f)

	log.Print("This is our first log message in Go!")



}