package util

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"net/http"
)

func Download(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	file, err := ioutil.TempFile("", "zipaddr_download_")
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		return "", err
	}

	return file.Name(), nil
}

func Unarchive(arcPath string) ([]string, error) {
	arcReader, err := zip.OpenReader(arcPath)
	if err != nil {
		return nil, err
	}
	defer arcReader.Close()

	paths := []string{}
	for _, file := range arcReader.File {
		fileReader, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer fileReader.Close()

		tmpFile, err := ioutil.TempFile("", "zipaddr.")
		if err != nil {
			return nil, err
		}
		defer tmpFile.Close()

		if _, err := io.Copy(tmpFile, fileReader); err != nil {
			return nil, err
		}

		paths = append(paths, tmpFile.Name())
	}

	return paths, nil
}

func UnarchiveURL(url string) ([]string, error) {
	path, err := Download(url)
	if err != nil {
		return nil, err
	}

	return Unarchive(path)
}
