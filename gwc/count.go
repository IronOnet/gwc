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

// TODO: Implement this function properly 
func processFiles(files []string) error{
	if len(files) == 0{
		fmt.Println("No files provided, use 'gwc' followed by file names")
		return nil 
	}

	for _, file := range files{
		data, err := readFile(file)
		if err != nil{
			fmt.Fprintf(os.Stderr, "Error reading file %q: %v\n", file, err)
			continue 
		}

		counts := count(data) 
		printResults(counts, file)
	}

	return nil 
}

func readFile(file string) ([]byte, error){
	return os.ReadFile(file) 
}

func count(data []byte) (map[string]int){
	counts := make(map[string]int) 

	counts["lines"] = 1 + strings.Count(string(data), "\n") 
	counts["words"] = countWords(data) 
	counts["chars"] = len(data) 
	counts["bytes"] = len(data) 

	return counts
}

func printResults(count map[string]int, file string){
	fmt.Println("Do nothing")
}

func countWords(data []byte) int{
	return len(string(data)) 
}
