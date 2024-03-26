package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Data struct {
	Name     string  `json:"name"`
	Language string  `json:"language"`
	Id       string  `json:"id"`
	Bio      string  `json:"bio"`
	Version  float32 `json:"version"`
}

type NewData struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Bio      string `json:"bio"`
}

func main() {
	jsonFileName := "test.json"

	// Read JSON file
	file, err := os.Open(jsonFileName)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
	}
	defer file.Close()

	// Read file contents
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal JSON data
	dataSlice := []Data{}
	if err := json.Unmarshal(buffer, &dataSlice); err != nil {
		fmt.Println("Error parsing JSON file:", err)
		return
	}

	// Create & Marshal New JsonData
	newJsonData := []NewData{}
	for _, data := range dataSlice {
		newJsonData = append(newJsonData, NewData{data.Name, data.Language, data.Bio})
	}
	jsonBytes, err := json.Marshal(newJsonData)
	if err != nil {
		fmt.Println("Error Marshal New Data:", err)
		return
	}
	fmt.Println("New Json Data:", string(jsonBytes))

	// Write New JsonData
	newJsonFileName := "new_test.json"
	newJsonFile, err := os.Create(newJsonFileName)
	if err != nil {
		fmt.Println("Error Create New file:", err)
		return
	}
	defer newJsonFile.Close()
	_, err = io.WriteString(newJsonFile, string(jsonBytes))
	if err != nil {
		fmt.Println("Error Write New file:", err)
		return
	}
}
