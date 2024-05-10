package osmodule

// https://www.honeybadger.io/blog/comprehensive-guide-to-file-operations-in-go/

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))
}

func CreateFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
}

func WriteToFile(filePath string, data string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	file.WriteString(data)
}

func MainFile() {
	filePath := "/tmp/test1.txt"
	// ReadFile(filePath)
	// CreateFile(filePath)
	file, err := os.Open(filePath)
	// to check if the error is of type file not exists
	if os.IsNotExist(err) {
		fmt.Println("File Not exists")
	}
	defer file.Close()
}
