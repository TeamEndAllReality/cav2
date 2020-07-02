package cav2

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writeStringToFile(str string, file string) error {
	return ioutil.WriteFile(file, []byte(str), os.ModePerm)
}

func readStringFromFile(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

//WriteBytesToFile Generic Write file bytes
func WriteBytesToFile(data []byte, file string) error {
	return ioutil.WriteFile(file, data, os.ModePerm)
}
