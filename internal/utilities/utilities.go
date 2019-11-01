package utilities

import (
	"os"
	"bufio"
)

// CheckStringInSlice checks if string s is in given array
func CheckStringInSlice(s string, arr []string) bool{
	for _, i := range arr{
		if i == s{
			return true
		}
	}
	return false
}

// PanicError if there is any kind of error panic
func PanicError(e error){
	if e != nil{
		panic(e)
	}
}

// ReturnFileLines gets fileLocation as argument and returns file line by line as a string slice
func ReturnFileLines(fileLocation string)(commandList []string){
	file, e := os.Open(fileLocation)
	PanicError(e)
	defer file.Close()

	fileLine := bufio.NewScanner(file)
	for fileLine.Scan(){
		commandList = append(commandList, fileLine.Text())
	}
	return commandList
}
