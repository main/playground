package main

import (
	"fmt"
	"os"
)

func main() {
	//parse args
	fmt.Printf("%#v", os.Args)
	if len(os.Args) != 3 {
		fmt.Println("Error: the number of command line parameters should be 2 (first one is source file path and second one is destination file path")
		os.Exit(1)
	}
	sourceFilePath := os.Args[1]
	destinationFilePath := os.Args[2]
	fmt.Printf("%#v%#v", sourceFilePath, destinationFilePath)
	//open source file
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer sourceFile.Close()
	//open destination file
	destinationFile, err := os.Create(destinationFilePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer destinationFile.Close()
	//TODO create variable for checksum
	//TODO start copying file
	//TODO start increment checksum variable

}
