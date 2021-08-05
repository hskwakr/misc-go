package main

import "github.com/hskwakr/misc-go/src/LoggingExamples/Tutorial_1/log/log"

func main() {
	//log.Println("This is a log message")
	//log.Fatal("This is a log message")
	//log.Panic("This is a log message")

	log.InfoLogger.Println("This is some info")
	log.WarningLogger.Println("This is probably important")
	log.ErrorLogger.Println("Something went wrong")
}
