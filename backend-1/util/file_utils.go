package util

import "os"

const FileName = "/home/C5311429/go/src/auto-test/record.json"

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
