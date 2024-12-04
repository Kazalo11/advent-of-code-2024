package util

import "os"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetData(fileName string) string {
	data, err := os.ReadFile(fileName)
	Check(err)
	return string(data)
}
