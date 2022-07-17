package gohelper

import (
	"io/ioutil"
	"os"
)

func makeDir(dirName string) error {
	return os.MkdirAll(dirName, 0755)
}

func writeFile(filename string, content []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func readFile(filename string) ([]byte, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return b, nil
}
