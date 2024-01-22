package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	WalkDir := "..\\Exercise10\\photos"
	filepath.Walk(WalkDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		curDir := filepath.Dir(path)
		m, err1 := rename(info.Name()) //get new name of the file
		if err1 != nil {
			fmt.Println(err1)
			return err1
		}
		oldPath := strings.Join([]string{curDir, info.Name()}, "\\") //creating old path
		newPath := strings.Join([]string{curDir, m}, "\\")           //creating new path
		err2 := os.Rename(oldPath, newPath)                          //changing file name
		if err2 != nil {
			fmt.Println(err2)
			return err2
		}
		return nil
	})
}

func rename(filename string) (string, error) {
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]              //getting file extention
	name := pieces[0]                         //getting name of file
	output := strings.FieldsFunc(name, Split) //splitting using brackets and spaces
	if len(output) <= 1 {
		return "", fmt.Errorf("string is not of valid format")
	}
	num := output[1]                                                           // getting file number after splitting the name
	number := strings.Join([]string{strings.Repeat("0", 3-len(num)), num}, "") //making new number by adding leading zeros
	newName := strings.Join([]string{output[0], number}, "_")                  //making new name using number
	newName = strings.Join([]string{newName, ext}, ".")                        //adding extension to the name
	return newName, nil
}
func Split(r rune) bool {
	return r == '(' || r == ' ' || r == ')'
}
