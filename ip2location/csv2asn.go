package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type ASNInfo struct {
	Name  string
	CIDRs []string
}

var asnMap = make(map[uint]ASNInfo)

func main() {
	files := []string{"IP2LOCATION-LITE-ASN.CSV"}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		reader := csv.NewReader(file)

		for {
			// 读取一行
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			if len(record) == 5 {
				if record[4] == "-" {
					continue
				}
				asnNum, err := strconv.ParseUint(record[3], 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				info, exists := asnMap[uint(asnNum)]
				if !exists {
					info = ASNInfo{
						Name:  record[4],
						CIDRs: []string{},
					}
				}
				info.CIDRs = append(info.CIDRs, record[2])
				asnMap[uint(asnNum)] = info
			}
		}
	}
	for asn, info := range asnMap {
		fmt.Println(asn, info)

	}
}
