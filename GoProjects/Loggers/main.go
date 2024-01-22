package main

import (
	"log"
	"os"
)

var infoLog, warnLog, debugLog, errorLog *log.Logger

func main() {
	file, error := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if error != nil {
		panic("Can not open file")
	}
	infoLog = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLog = log.New(file, "WARN ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLog = log.New(file, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(file, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	nos := []int{-1, 0, 2, 3}
	findEven(nos)
}

func findEven(nos []int) {
	infoLog.Println("INput array is ", nos)
	for _, no := range nos {
		debugLog.Printf("Current no. is %d", no)
		if no < 0 {
			errorLog.Println(no, " is less than 0")
			continue
		}
		if no == 0 {
			warnLog.Printf("No is equal to zero")
		}
		if no%2 == 0 {
			infoLog.Printf("%d is even no.", no)
		} else {
			infoLog.Printf("%d is odd no.", no)
		}
	}
}

//logrus, third party logging package, tdt-table driven testing, golang benchmarking,how to execute specific test case and specific test file from terminal, go test and go tool to write coverage on a html file
