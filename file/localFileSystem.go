package file

import ("fmt"
		"os"
		"bufio"
		"strings"
		"errors")


func GetNewLocalFileSystem(filePath string) (LocalFileSystem, error){
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0644)
	if(err!=nil){
		return LocalFileSystem{} , err
	}
	return LocalFileSystem{"myFileSystem",file}, nil
}

type LocalFileSystem struct{
	Name string
	File *os.File
}