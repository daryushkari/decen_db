package utilities

import (
	"bufio"
	"os"
)

// CheckStringInSlice checks if string s is in given array
func CheckStringInSlice(s string, arr []string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}

// ReturnFileLines gets fileLocation as argument and returns file line by line as a string slice
func ReturnFileLines(filePath string) (linesList []string, err error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileLine := bufio.NewScanner(file)
	for fileLine.Scan() {
		linesList = append(linesList, fileLine.Text())
	}
	return linesList, err
}

//AppendFile gets lines as array of strings and appends to file given
func AppendFile(lines []string, filePath string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)

	if err != nil {
		return err
	}

	defer f.Close()
	for _, i := range lines {
		if _, err := f.WriteString(i + "\n"); err != nil {
			if err != nil {
				return err
			}
		}
	}

	return nil

}
