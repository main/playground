package main

import (
	"fmt"
	"os"
)

func main() {
	//TODO parse args
	fmt.Printf("%#v", os.Args)
	if len(os.Args) != 3 {
		fmt.Println("Error: the number of command line parameters should be 2 (first one is source file path and second one is destination file path")
		os.Exit(1)
	}
	sourceFilePath := os.Args[1]
	destinationFilePath := os.Args[2]
	fmt.Printf("%#v%#v", sourceFilePath, destinationFilePath)
	//TODO open source file
	file, err := os.Open(sourceFilePath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	//TODO open destination file
	//TODO create variable for checksum
	//TODO start copying file
	//TODO start increment checksum variable

}
