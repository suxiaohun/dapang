package aliyun

import (
	"encoding/csv"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"os"
)

func GetObjectToFile() {
	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", "your ak", "your sk")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	bucket, err := client.Bucket("your bucket name")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	csvFile, err := os.Open("oss/aliyun/test.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	indexSub := 0
	for i, line := range csvLines {
		if i == 0 {
			continue
		}
		fileName := line[1]
		indexSub++
		err = bucket.GetObjectToFile(fileName, "oss/aliyun/"+fileName+".jpg")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		fmt.Println(indexSub, fileName)
	}
	os.Exit(0)
}

func mergeCsvFiles() {
	csvFile1, err := os.Open("test1.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile1.Close()

	csvLines1, err := csv.NewReader(csvFile1).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	data := make(map[string]string, 0)
	for i, line := range csvLines1 {
		if i == 0 {
			continue
		}
		requestID := line[2]
		data[requestID] = line[1]
	}

	csvFile2, err := os.Open("test2.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile2.Close()

	csvLines2, err := csv.NewReader(csvFile2).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("merge.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to create file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for i, line := range csvLines2 {
		if i == 0 {
			continue
		}

		record := []string{line[2], data[line[2]]}
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
