package main

import (
	"fmt"
	"os"

	"github.com/irononet/gwc/gwc"
	"github.com/spf13/cobra"
)

var version string = "0.1.0"

var description string = `
GWC(1)		User Commands		GWC(1)

NAME
	wc - print newline, word, and byte counts for each file

SYNOPSIS 
	wc [OPTION]... [FILE]... 
	wc [OPTION]... --files0-from=F 

DESCRIPTION 
	Print newline, word and byte counts for each FILE, and and total line if more than
	one FILE is specified. A word is non-zero-length sequence of characters delimited by 
	white space.

	With no FILE, or when FILE is -, read standard input. 

	The options below may be used to select which counts are printed, always in the following order: 
	newline, word, character, byte, maximum line length.

	-c, --bytes
		print the bytes counts 

	-m, --chars 
		print the character counts 

	-l, --lines 
		print the newline counts 

	-files0-fro=E 
		read input from the files specified by NUL-terminated names in file F, IF F is - then read names 
		from standard input

	-L, --max-line-length 
		print the maximum display width 

	-w, --words 
		print the word counts 

	--help display this help and exit 

	--version 
		output the version information and exit 

AUTHOR
	Written by Arnaud Wanet (IronOnet)

`


var mockCmd *cobra.Command = &cobra.Command{
	Use: "mock", 
	Short: "mock command, that does nothing", 
	Long: "This command does nothing", 
	Run: mockCommand,
}


var rootCmd *cobra.Command = &cobra.Command{
	Use: "gwc", 
	Short: "gwc, print newline, word, and bytes counts for each file", 
	Long: description, 
	Run: rootCmdF, // Prevents default behavior
}

func init(){

	rootCmd.Flags().BoolP("chars", "m", false, "count the number of characters in a file")
	rootCmd.Flags().BoolP("bytes", "c", false, "count the number of bytes in a file") 
	rootCmd.Flags().BoolP("words", "w", false, "count the number of words in a file") 
	rootCmd.Flags().BoolP("lines", "l", false, "count the number of lines in a file") 
	rootCmd.Flags().BoolP("version", "v", false, "display the version number")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	charFlag, err  := rootCmd.Flags().GetBool("chars")
	if err != nil{
		fmt.Println(err) 
	}
	bytesFlag, err := rootCmd.Flags().GetBool("bytes") 
	if err != nil{
		fmt.Println(err) 
	}
	wordsFlag, err := rootCmd.Flags().GetBool("words") 
	if err != nil{
		fmt.Println(err) 
	}
	linesFlag, err := rootCmd.Flags().GetBool("lines") 
	if err != nil{
		fmt.Println(err) 
	}

	versionFlag, err := rootCmd.Flags().GetBool("version") 
	if err != nil{
		fmt.Println(err)
	}

	if charFlag{
		if len(os.Args) > 2{
			filepath := os.Args[2] 
			res, err := gwc.CountChar(filepath) 
			if err != nil{
				printFileNotFoundError(filepath)
			}
			fmt.Printf("%d %s\n", res, filepath)
		} else{
			fmt.Println("No file argument provided.") 
		}
	}
	if bytesFlag{
		if len(os.Args) > 2{
			filepath := os.Args[2] 
			res, err := gwc.CountBytes(filepath) 
			if err != nil{
				fmt.Printf("file not found %s\n", filepath) 
			}
			fmt.Printf("%d %s\n", res, filepath) 
		} else{
			fmt.Println("No file argument provided.") 
		}
	}

	if wordsFlag{
		if len(os.Args) > 2{
			filepath := os.Args[2] 
			res, err := gwc.CountWords(filepath) 
			if err != nil{
				printFileNotFoundError(filepath)
			}
			fmt.Printf("%d %s\n", res, filepath) 
		} else{
			fmt.Println("No file argument provided.") 
		}
	}

	if linesFlag{
		if len(os.Args) >2{
			filepath := os.Args[2] 
			res, err := gwc.CountLines(filepath) 
			if err != nil{
				printFileNotFoundError(filepath)
			}
			fmt.Printf("%d %s\n", res, filepath) 
		} else{
			fmt.Println("No file argument passed") 
		}
	}

	if versionFlag{
		fmt.Println("gwc version " + version) 
	}

	if len(os.Args) >= 1 && (!linesFlag) && (!bytesFlag) && (!wordsFlag) && (!charFlag){
		filepath := os.Args[1] 
		resWords, _ := gwc.CountWords(filepath) 
		resBytes, _ := gwc.CountBytes(filepath) 
		resLines, _ := gwc.CountLines(filepath) 
		resChars, _ := gwc.CountChar(filepath) 
		fmt.Printf("%d %d %d %d %s \n", resWords, resBytes, resLines, resChars, filepath) 
	}


}

func rootCmdF(cmd *cobra.Command, args []string){

}

func printFileNotFoundError(filename string){
	fmt.Printf("error file %s not found\n", filename)
}

func mockCommand(cmd *cobra.Command, args []string){
	fmt.Println("this is a mock command")
}