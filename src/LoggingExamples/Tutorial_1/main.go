package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func main() {
	//log.Println("This is a log message")
	//log.Fatal("This is a log message")
	//log.Panic("This is a log message")

	InfoLogger.Println("This is some info")
	WarningLogger.Println("This is probably important")
	ErrorLogger.Println("Something went wrong")
}
