package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var pathToSecondArg string
	arguments := os.Args // type []string
	// Check if file or folder

	switch len(arguments){
	case 2:
		if arguments[1] == "-h" || arguments[1] == "--help" {
			printHelp()
			return
		}

	case 3:
		if !strings.Contains(arguments[2], "/") {
			pathToSecondArg = getLocation(arguments)
			e := os.Rename(arguments[1], pathToSecondArg+arguments[2])
			if e != nil {
				fmt.Println(e)
			}
		} else {
			fmt.Println("Invalid Argument: This program assumes" +
				" the path.")
		}
	case 4:
		if arguments[1] == "-f" {
			fmt.Println ("Some folder-y stuff")
		}
	default:
		fmt.Println("Invalid Argument" +
			" Check Usage via \"rname -h\"")
	}
}

func getLocation([]string arguments) string {
	if strings.Contains(arguments[1],"/") {
		str := arguments [1]
		s := strings.Split(str, "/")
		return str[:len(str) - len(s[len(s)-1])]
	} else {
		return ""
	}
}

func printHelp() {
	fmt.Println("Usage:\n\t" +
		"rname [path to file] [new name]\n\n" +
		"Arguments:\n\t" +
		"-f\t\tchange the name of a folder\n\t" +
		"-h, --help\tshows usage and arguments\n\t")
}
