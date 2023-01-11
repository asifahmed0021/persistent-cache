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

func (f LocalFileSystem) Append(key string, value string) error{
	fileScanner := bufio.NewScanner(f.File)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
        text := fileScanner.Text()
		words := strings.Fields(text)
		if(words[0]==key){
			return errors.New("key already exists")
		}
    }

	f.File.WriteString(key + " " + value + "\n")

	return nil
}

func (f LocalFileSystem) GetValueForKey(key string) (string, error){
	readFile := f.File
  
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
        text := fileScanner.Text()
		words := strings.Fields(text)
		if(words[0]==key){
			return words[1],nil
		}
    }

	return  "", errors.New("no key found")
}