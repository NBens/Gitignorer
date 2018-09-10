package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// GitignoreFiles : Url to the zipped gitignore files
const GitignoreFiles = "https://github.com/github/gitignore/archive/master.zip"

// DownloadFile : Downloads a files from an url, and saves it to a path
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

// UnzipFile : Unzips a zip file in its current folder
func UnzipFile(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			os.MkdirAll(fpath, os.ModePerm)

		} else {

			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			outFile.Close()

			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}

// FilesNamesDir : Return an array with files name without the extension
func FilesNamesDir(filepath, extension string) ([]string, error) {
	files := []string{}
	data, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, err
	}
	for _, v := range data {
		fileName := v.Name()
		if strings.HasSuffix(fileName, extension) {
			file := fileName[:len(fileName)-len(extension)]
			files = append(files, file)
		}
	}
	return files, nil
}

// list : Lists the available languages' gitignore files, global gitignore files, and templates
func list() {

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

// update : Update command, downloads gitignore files from github, extracts them to gitignorer_data
func update() {

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

// showHelp : Shows the available commands
func showHelp() {
	fmt.Println(`
Available commands:
===================
-list   : Lists the available languages' gitignore files, global gitignore files, and templates
-update : Update command, downloads gitignore files from github, extracts them to gitignorer_data
				`)
}

func main() {

	if len(os.Args) == 1 {
		showHelp()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "update":
		update()
	case "create":
		fmt.Println("Create")
	case "list":
		list()
	case "create-template":
		fmt.Println("Create Template")
	case "list-templates":
		fmt.Println("List Templates")
	case "use-template":
		fmt.Println("Use Template")
	default:
		showHelp()
	}

}
