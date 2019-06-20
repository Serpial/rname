# rname
## Reasoning
I know you can rename files with 'mv' but I wanted to make this project to practise Go and make it a little easier to rename files.
## Features
**Allows the user to change the name of an individual file**
* Constricts the user to using a whitelist of character for naming the new file
* Validates that the file that they wish to change exists
* Validates that there are no files in the directory that are the same as the new name
## Usage
At the moment, the project can only deal with renaming individual files. To use **rname**, the first argument of the command should be location of the file (This can be the current directory). The second argument should be the file name that you wish to replace it with. <br>
`rname <Location of file> <New file name>`<br>
For example,<br>
`rname ~/Documents/rnm.go rname.go`
## Installation
Until I fix up the Make file, the current way of installation would require the user to install Go and run `go build rname.go` in rname's development directory. Then move that file into your bin so that you can use it anywhere.
## In the future
* I would like to flesh out the Make file for testing and building purposes
