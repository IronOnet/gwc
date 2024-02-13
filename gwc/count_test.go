package gwc

import (
	"testing"
)



func TestGWC_CountBytes(t *testing.T) {

	t.Run("case where the file is empty", func(t *testing.T) {
		filepath := "test_empty.txt"
		count, err := CountBytes(filepath)

		// An error should be raised
		if err != nil {
			t.Errorf("For an empty file an error should be raised")
		}

		if count != 0 {
			t.Errorf("For an empty file the count should be equal to zero")
		}
	})

	t.Run("case where the file is non empty", func(t *testing.T) {
		filepath := "test.txt"
		byteCount, err := CountBytes(filepath)

		if err != nil {
			t.Errorf("An error occurred %s", err)
		}

		if !(byteCount > 0) {
			t.Errorf("byte count should be greater than zero for non empty file")
		}
	})

	t.Run("test for test.txt", func(t *testing.T) {
		filepath := "test.txt"
		actualByteCount, err := CountBytes(filepath)

		if err != nil {
			t.Errorf("An error occurred: %s", err)
		}

		var expectedByteCount int = 342190

		if actualByteCount != expectedByteCount {
			t.Errorf("for file %s expected a byte count of %d but got %d", filepath, expectedByteCount, actualByteCount)
		}
	})

}

func TestGWC_CountLines(t *testing.T) {
	t.Run("returns 0 if file is empty", func(t *testing.T) {
		filepath := "test_empty.txt"
		expectedCount := 0

		actualCount, err := CountLines(filepath)
		if err != nil {
			t.Errorf("an error occurred: %s", err)
		}

		if actualCount != expectedCount {
			t.Errorf("for empty file %s expected line count to be %d but got %d", filepath, expectedCount, actualCount)
		}

	})

	t.Run("returns the expected number of lines", func(t *testing.T) {
		filepath := "test.txt"

		expectedCount := 7145

		actualCount, err := CountLines(filepath)
		if err != nil {
			t.Errorf("an error occurred: %s", err)
		}

		if actualCount != expectedCount {
			t.Errorf("for file %s expected count %d but got %d", filepath, expectedCount, actualCount)
		}
	})
}

func TestGWC_CountWords(t *testing.T) {
	t.Run("for an empty file return 0", func(t *testing.T){
		filepath := "test_empty.txt" 

		expectedCount := 0 
		actualCount, err := CountWords(filepath) 
		if err != nil{
			t.Errorf("an error occurred: %s", err) 
		}

		if expectedCount != actualCount{
			t.Errorf("for empty file %s expected count %d but got %d", filepath, expectedCount, actualCount)
		}
	})

	t.Run("for test.txt", func(t *testing.T){
		filepath := "test.txt" 

		expectedCount := 58164 

		actualCount, err := CountWords(filepath) 
		if err != nil{
			t.Errorf("an error occurred: %s", err) 
		}

		if actualCount != expectedCount{
			t.Errorf("for file %s expected word count %d but got %d", filepath, expectedCount, actualCount)
		}
	})

}

func TestGWC_CountChars(t *testing.T) {
	t.Run("empty file return 0 char count", func(t *testing.T){
		filepath := "test_empty.txt" 

		expectedCount := 0 

		actualCount, err := CountChar(filepath) 
		if err != nil{
			t.Errorf("an error occurred: %s", err) 
		}

		if expectedCount != actualCount{
			t.Errorf("for file %s expected count %d but got %d", filepath, expectedCount, actualCount)
		}
	})

	t.Run("count chars for test.txt", func(t *testing.T){
		filepath := "test.txt" 

		expectedCount := 342190

		actualCount, err := CountChar(filepath) 
		if err != nil{
			t.Errorf("an error occurred: %s", err) 
		}

		if actualCount != expectedCount{
			t.Errorf("for file %s expected char count %d but got %d", filepath, expectedCount, actualCount)
		}
	})
}

func Benchmark_CountBytes(b *testing.B){

}

func Benchmark_CountWords(b *testing.B){

}

func Benchmark_CountChars(b *testing.B){

}

func Benchmark_CountLines(b *testing.B){

}
