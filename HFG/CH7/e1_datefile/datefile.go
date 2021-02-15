package e1_datefile

import (
	"bufio"
	"os"
)

// getstrings reads a string each line of file.
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil{
		return nil, scanner.Err()
	}
	return lines, nil
}

