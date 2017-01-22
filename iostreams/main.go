package main

import (
	"fmt"
	"io"
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

	//start copying file
	bufer := [1 << 10]byte{}
	for {
		n, err := sourceFile.Read(bufer[:])
		if n == 0 && err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			panic(err)
		}

		n, err = destinationFile.Write(bufer[:])
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	//TODO start increment checksum variable

}
