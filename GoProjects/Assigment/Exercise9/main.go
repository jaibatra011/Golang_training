package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

var myFile *os.File

func walk(s string, d fs.DirEntry, err error) error {
	file, err1 := os.OpenFile(myFile.Name(), os.O_APPEND|os.O_WRONLY, 0666)

	if err1 != nil {
		panic(err1)
	}

	defer file.Close()

	if !d.IsDir() {
		_, err2 := file.WriteString(s + " \n")
		if err2 != nil {
			panic(err2)
		}
	}
	return nil
}

func main() {

	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err.Error())
	}

	myFile = file

	filepath.WalkDir("..\\Exercise7", walk)
}
