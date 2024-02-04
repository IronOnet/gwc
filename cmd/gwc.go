package main

import "fmt"

var description string = `
	GWC(1)	User Commands		GWC(1)

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

func main() {
	fmt.Println(description)
}