package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args // type []string

	switch len(arguments){
	case 2:
		if arguments[1] == "-h" || arguments[1] == "--help" {
			printHelp()
			return
		}
	case 3:
		if isFile(arguments) == 1 {
			if !strings.Contains(arguments[2], "/") {
				doRename(arguments)
			} else {
				fmt.Println("Invalid Argument: This program assumes" +
					" the path.")
			}
		} else {
			basicError()
		}
	case 4:
		if isFile(arguments) == 2 {
			if arguments[1] == "-f" {
				fmt.Println ("Some folder-y stuff")
			}
		} else {
			basicError()
		}
	default:
		basicError()
	}
}

func getLocation(arguments []string) string {
	if strings.Contains(arguments[len(arguments)-2],"/") {
		str := arguments[len(arguments)-2]
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

func isFile (arguments []string) byte {
	fi, err := os.Stat(getLocation(arguments)+arguments[len(arguments)-2])
    if err != nil {
        fmt.Println(err)
        return 0
    }
	
    switch mode := fi.Mode(); {
	case mode.IsRegular():
		return 1
    case mode.IsDir():
		return 2
    }
	return 0
}

func basicError() {
	fmt.Println("Invalid Argument:" +
		" Check Usage via \"rname -h\"")
}

func doRename(arguments []string) {
	newNameAndLoc := getLocation(arguments)+arguments[len(arguments)-1]
	e := os.Rename(arguments[len(arguments)-2], newNameAndLoc)
	if e != nil {
		fmt.Println(e)
	}
}
