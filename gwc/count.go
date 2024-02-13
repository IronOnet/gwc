package gwc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CountBytes(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, fmt.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return 0, fmt.Errorf("Error getting file info: %s", err)
	}

	byteCount := fileInfo.Size()

	return int(byteCount), nil
}

func CountChar(filepath string) (int, error) {

	content , err := os.ReadFile(filepath) 
	if err != nil{
		return 0, fmt.Errorf("could not open file an error occurred: %s", err) 
	} 
	characterCount := len(content) 
	

	return characterCount, nil 

}

func CountLines(filepath string) (int, error) {
	file, err := os.Open(filepath) 
	if err != nil{
		return 0, fmt.Errorf("could not open file %s an error occured %s", filepath, err) 
	}
	defer file.Close() 

	scanner := bufio.NewScanner(file) 
	// Set split function to identify lines as separator 
	scanner.Split(bufio.ScanLines)  

	lineCount := 0 
	if err = scanner.Err(); err != nil{
		return 0, fmt.Errorf("error scanning file: %s", err) 
	}

	for scanner.Scan(){
		lineCount++ 
	}

	return lineCount, nil 
}

func CountWords(filepath string) (int, error) {
	file, err := os.Open(filepath) 
	if err != nil{
		return 0, fmt.Errorf("could not open file, an error occurred: %v", err) 
	}
	defer file.Close() 

	scanner := bufio.NewScanner(file) 

	if err := scanner.Err(); err != nil{
		return 0, fmt.Errorf("error scanning file: %s", err) 
	}

	wordCount := 0 
	for scanner.Scan(){
		words := strings.Fields(scanner.Text()) 
		wordCount += len(words) 
	}

	return wordCount, nil 
}
