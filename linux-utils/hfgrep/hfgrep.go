package main

import (
	"fmt"
	//"bufio"
	"os"
)

func main() {
	fmt.Println("Testeys")

	//if no args are passed
	if len(os.Args) < 2 {
			fmt.Println("hfgrep: searchterm [file ...]")
			os.Exit(1)
	}


	fmt.Println("uhhhh")
	
}
