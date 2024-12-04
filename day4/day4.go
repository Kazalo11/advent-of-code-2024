package main

import "os"

func getData() string {
	data, err := os.ReadFile("day3.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {

	data := getData("day4.txt")

}
