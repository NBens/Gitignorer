package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
)

// HomeDir : Extracts the home directory
func HomeDir() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(currentUser.HomeDir)
	}
}

// DownloadFile : Downloads a files from an url, and saving it to a path
func DownloadFile(path, url string) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := http.Get(url)
	if err != nil {
		return err
	}
	defer data.Body.Close()

	_, err = io.Copy(file, data.Body)
	if err != nil {
		return err
	}

	return nil

}

func main() {

}
