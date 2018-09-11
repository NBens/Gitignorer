package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		ShowHelp()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "update":
		Update()
	case "create":
		if len(os.Args) > 2 && strings.TrimSpace(os.Args[2]) != "" && IsFileExist("./gitignorer_data") {
			Create(os.Args[2])
		} else {
			ShowHelp()
			os.Exit(1)
		}
	case "list":
		List()
	case "create-template":
		fmt.Println("Create Template")
	case "list-templates":
		fmt.Println("List Templates")
	case "use-template":
		fmt.Println("Use Template")
	default:
		ShowHelp()
	}

}
