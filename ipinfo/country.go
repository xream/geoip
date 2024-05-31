package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 打开 CSV 文件
	file, err := os.Open("./country.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建 CSV 阅读器
	reader := csv.NewReader(file)
	
	// 读取 CSV 文件头
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading headers:", err)
		return
	}

	// 检查必要的列索引
	rangeIdx := -1
	countryIdx := -1
	for i, header := range headers {
		switch header {
		case "range":
			rangeIdx = i
		case "country":
			countryIdx = i
		}
	}

	if rangeIdx == -1 || countryIdx == -1 {
		fmt.Println("Required columns not found")
		return
	}

	// 用于存储每个国家对应的 IP 范围
	countryData := make(map[string][]string)

	// 读取 CSV 文件内容
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading record:", err)
			return
		}

		ipRange := record[rangeIdx]
		country := strings.ToLower(record[countryIdx])

		// 将 IP 范围加入到对应国家的列表中
		countryData[country] = append(countryData[country], ipRange)
	}

	// 创建 data 目录
	err = os.MkdirAll("../data", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating data directory:", err)
		return
	}

	// 将每个国家的数据写入到对应的文件中
	for country, ranges := range countryData {
		filePath := filepath.Join("../data", country)
		content := strings.Join(ranges, "\n")
		err := writeToFile(filePath, content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
}

func writeToFile(filePath string, content string) error {
	// 创建或打开文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 一次性写入文件
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
