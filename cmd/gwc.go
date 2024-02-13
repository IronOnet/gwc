package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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

var rootCmd *cobra.Command = &cobra.Command{
	Use:   "gwc",
	Short: "gwc, print newline, word, and bytes counts for each file",
	Long:  description,
}

var mockCmd *cobra.Command = &cobra.Command{
	Use: "mock", 
	Short: "mock command, that does nothing", 
	Long: "This command does nothing", 
	Run: mockCommand,
}

func main() {
	rootCmd.AddCommand(mockCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func mockCommand(cmd *cobra.Command, args []string){
	fmt.Println("this is a mock command")
}

func countCharsCmdF(cmd *cobra.Command, args []string){

}

func countWordsCmdF(cmd *cobra.Command, args []string){

}

func countLinesCmdF(cmd *cobra.Command, args []string){

}

func countBytesCmdF(cmd *cobra.Command, args []string){
	
}
