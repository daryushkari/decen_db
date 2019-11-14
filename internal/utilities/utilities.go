package utilities

import (
	"bufio"
	"os"
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
func PanicError(err error){
	if err != nil{
		panic(err)
	}
}

// ReturnFileLines gets fileLocation as argument and returns file line by line as a string slice
func ReturnFileLines(filePath string)(linesList []string){
	file, err := os.Open(filePath)
	PanicError(err)
	defer file.Close()

	fileLine := bufio.NewScanner(file)
	for fileLine.Scan(){
		linesList = append(linesList, fileLine.Text())
	}
	return linesList
}

func AppendFile(lines []string, filePath string){
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)
	PanicError(err)
	defer f.Close()
	for _, i := range lines{
		if _, err := f.WriteString(i + "\n"); err != nil {
			PanicError(err)
		}
	}

}
