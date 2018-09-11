package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// GitignoreFiles : Url to the zipped gitignore files
const GitignoreFiles = "https://github.com/github/gitignore/archive/master.zip"

// Create : Create command, Creates gitignore files from a list of languages/globals
func Create(languages string) {
	filepath := "./gitignorer_data/gitignores/"
	globalPath := filepath + "Global/"
	languageSlice := strings.Split(languages, ",")
	outData := []byte{}
	for _, v := range languageSlice {
		fullName := v + ".gitignore"
		if !IsFileExist(filepath+fullName) && !IsFileExist(globalPath+fullName) {
			fmt.Println("Language/global gitignore " + v + " does not exist, skipping it")
			continue
		} else if IsFileExist(filepath + fullName) {
			fullName = filepath + fullName
		} else if IsFileExist(globalPath + fullName) {
			fullName = globalPath + fullName
		}
		gitignoreData, err := ReadFile(fullName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		outData = append(outData, []byte("\n##### "+v+" #####\n\n")...)
		outData = append(outData, gitignoreData...)
	}
	mode := int(0777)
	ioutil.WriteFile(".gitignore", outData, os.FileMode(mode))
}

// List : Lists the available languages' gitignore files, global gitignore files, and templates
func List() {

	fmt.Println("\nList of available languages:")
	fmt.Println("============================")
	languages, err := FilesNamesDir("./gitignorer_data/gitignores", ".gitignore")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(strings.Join(languages, ", "))

	fmt.Println("\nGobal useful Gitignores:")
	fmt.Println("========================")
	globals, err := FilesNamesDir("./gitignorer_data/gitignores/Global", ".gitignore")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(strings.Join(globals, ", "))

	fmt.Println("\nAvailable Templates:")
	fmt.Println("====================")
	templates, err := FilesNamesDir("./gitignorer_data/Templates", ".Template.gitignore")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(strings.Join(templates, ", "))

}

// Update : Update command, downloads gitignore files from github, extracts them to gitignorer_data
func Update() {

	fmt.Println("Update Gitignore Files")

	fmt.Printf("Downloading Gitignore files...")
	DownloadFile("gitignores.zip", GitignoreFiles)
	fmt.Printf(" Done\n")

	fmt.Printf("Unzipping files...")
	_, err := UnzipFile("gitignores.zip", "./gitignorer_data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(" Done\n")

	os.Remove("gitignores.zip")

	rename := os.Rename("./gitignorer_data/gitignore-master", "./gitignorer_data/gitignores")
	if rename != nil {
		fmt.Println(err)
	}

	os.Mkdir("./gitignorer_data/Templates", os.ModePerm)

	fmt.Println("Updating Done!")
}

// ShowHelp : Shows the available commands
func ShowHelp() {
	fmt.Println(`
Available commands:
===================
-update : Downloads gitignore files from github, extracts them to gitignorer_data
-list   : Lists the available languages' gitignore files, global gitignore files, and templates
-create : Creates gitignore files from a list of languages/globals sparated by commas. Example: "create python,java,emacs"
				`)
}
