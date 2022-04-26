package main

import (
	"fmt"
	"log"
	"go.uber.org/zap"
)

func main(){
	fmt.Println("Hello, world!")

	// Firstly,we create a configuration (i.e.  builder) for our logger
	config := zap.NewProductionConfig()

	// we configure the destination for our log, in this case, a file

	config.OutputPaths = []string{"zap-demo.log"}

	// we build the logger

	zapLogger, err := config.Build()


	// Go doesn't support exceptions, so we check for the error, exiting the application if needed

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err);
	}

	// flushes buffer, if any

	defer zapLogger.Sync()

	// We finally log our message at the INFO level

	zapLogger.Info("This is our first log message using zap!")


	// f, err := os.OpenFile("golang-demo.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }

	// defer f.Close()

	// log.SetOutput(f)

	//log.Print("This is our first log message in Go!")

}