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

// ReadFile : Returns file contents as a string
func ReadFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, err
}

// IsFileExist : Checks if a file/directory does exist or not
func IsFileExist(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}
