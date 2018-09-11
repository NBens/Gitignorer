package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		ShowHelp()
	}
	switch os.Args[1] {
	case "update":
		Update()
	case "create":
		if len(os.Args) > 2 && strings.TrimSpace(os.Args[2]) != "" && IsFileExist("./gitignorer_data") {
			Create(os.Args[2], ".gitignore")
		} else {
			ShowHelp()
		}
	case "list":
		List()
	case "create-template":
		if len(os.Args) > 3 && strings.TrimSpace(os.Args[2]) != "" && IsFileExist("./gitignorer_data") {
			Create(os.Args[2], "./gitignorer_data/Templates/"+os.Args[3]+".Template.gitignore")
		} else {
			ShowHelp()
		}
	case "use-template":
		if len(os.Args) > 2 && strings.TrimSpace(os.Args[2]) != "" && IsFileExist("./gitignorer_data") {
			if IsFileExist("./gitignorer_data/Templates/" + os.Args[2] + ".Template.gitignore") {
				useTemp := UseTemplate(os.Args[2], "Template."+os.Args[2]+".gitignore")
				if useTemp != nil {
					fmt.Println(useTemp)
				} else {
					fmt.Println("Saved as: " + "Template." + os.Args[2] + ".gitignore")
				}
			} else {
				fmt.Println("Template doesn't exist")
			}
		} else {
			ShowHelp()
		}
	default:
		ShowHelp()
	}

}
