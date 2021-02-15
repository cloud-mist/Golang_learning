// Package e5_datefile allows reading date samples from files.
package e5_datefile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file.
// 将返回一个数字数组和遇到的错误
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	// 如果关闭和扫描文件出现错误，就返回错误
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	// 如果没有错误，就返回数字数组和nil错误
	return numbers, nil
		
}
