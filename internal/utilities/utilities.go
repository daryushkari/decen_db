package utilities

import (
	"bufio"
	"os"
)

// ReturnFileLines gets fileLocation as argument and returns file line by line as a string slice
func ReturnFileLines(filePath string) (linesList []string, err error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer CloseFile(file, &err)

	fileLine := bufio.NewScanner(file)
	for fileLine.Scan() {
		linesList = append(linesList, fileLine.Text())
	}
	return linesList, err
}

//AppendFile gets lines as array of strings and appends to file given
func AppendFile(lines []string, filePath string) (err error) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0700)

	if err != nil {
		return err
	}

	defer CloseFile(f, &err)

	for _, i := range lines {
		if _, err := f.WriteString(i + "\n"); err != nil {
			return err
		}
	}

	return nil

}
