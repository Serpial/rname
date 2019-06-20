package main

import (
	"fmt"
	"os"
	"strings"
)

/*
* The main function is just the skeleton of what the program has to do.
* First it checks the arguments that the user has entered, then it 
* renames the file based on those arguments
*/
func main() {
	arguments := os.Args
	
	if checkInput(arguments) == 0 {
		renameFile(arguments)
	}
}


/*
* Returns 0 if the input is fine and we are good to do a rename
* Returns 1 if there is an issue with the input
* Returns 2 if the input was fine and we just returned the usage screen
*/
func checkInput(arguments []string) int {
	if len(arguments)<2 || len(arguments)>3{
		fmt.Println("rname: There was an issue with the" +
			" number of arguments you entered")
		return 1
	} else if len(arguments)==2 {
		// This is where other arguements can be selected like '--help'
		if arguments[1]=="--help" || arguments[1]=="-h" {
			printHelpAndUsage()
		} else {
			fmt.Println("Not an option")
			fmt.Println("See -h or --help for information")
		}
		return 2
	}

	// Check to see if the first argument is a file that exists
	valid := checkIfValidFile(arguments[1])
	if valid>0 {
		if valid==1 {
			fmt.Println("Error: This program cannot be used with directories : "+
				arguments[1])			
		} else {
			fmt.Println("Error: No such file or directory : " + arguments[1])
		}
		return 1
	}

	// Now that we have dealt with the 1st argument, it would be a good idea
	//    to sanitize the second argument
	if sanitizeNewFileName(arguments)==1 { return 1	}

	// Check if the new name already exists
	valid = checkIfValidFile(getLocation(arguments[1])+arguments[2])
	if valid<2 {
		if valid==0 {
			fmt.Println("Error: New file name already exists in this directory")
		} else {
			fmt.Println("Error: New file name is the same as a directory")
		}
		return 1
	}
	
	return 0
}


/*
* Gets the file path of the file the user wishes to change
*/
func getLocation(originalFile string) string {
	ind := strings.LastIndex(originalFile, "/");
	if  ind != -1{
		// Return the whole directory path
		return originalFile[:ind+1]
	} else {
		// If the original file location does not contain a "/" then assume
		//    that it is in the current directory
		return "./"
	}
}


/*
* Uses whitelisting to restrict the characters the user can use 
*/
func sanitizeNewFileName(arguments []string) int {
	whiteList := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890" +
		",.@#[]{}¬`!£$%^()_-+= " 

	// I read this was a standard character limit for linux file
	if len(arguments[2])>255 {
		fmt.Println("Error: New name is too large")
	}

	// Compare each character against the whitelist to see if each letter is on it
	for _, el := range arguments[2] {
		if !(strings.Contains(whiteList, string(el))) {
			fmt.Println("Error: New name contains illicit character" +
				" : \"" + string(el) + "\"")
			return 1;
		}
	}
	return 0
}


/*
* To be used to check if the file is a file and if it already exists
* Returns 0 if the file exists
* Returns 1 if the file is actually a directory
* Returns 2 if the file does not exist
*/
func checkIfValidFile(fileString string) int {
	file, err := os.Stat(fileString)
	if err != nil {	return 2 }

	if file.Mode().IsDir() {
		return 1
	}
	return 0
}


/*
* This prints all of the stuff I want printed in the help section
*/
func printHelpAndUsage() {
	fmt.Println("Basic program that allowes the renaming" +
		" of an individual file")
	fmt.Println("Usage: rname <File Location> <New Name>" +
		"\n\t e.g. rname ~/Documents/file fileTwo")
	fmt.Println("\nOptions:\n\t--help or -h\tDisplay usage" +
		" and options")
}


/*
* This function is the function that actually renames the file and gives out
* an error if it gets any
*/
func renameFile(arguments []string) {
	e := os.Rename(arguments[1], getLocation(arguments[1])+ arguments[2])
	if e != nil {
		fmt.Println(e)
	}
}
