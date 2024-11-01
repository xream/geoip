package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
)

type ASNInfo struct {
	Name  string
	CIDRs []string
}

var asnMap = make(map[uint]ASNInfo)

func main() {
	files := []string{"IP2LOCATION-LITE-ASN.CSV", "IP2LOCATION-LITE-ASN.IPV6.CSV"}

	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			log.Fatalln(err)
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
				log.Fatalln(err)
			}
			if len(record) == 5 {
				if record[4] == "-" {
					continue
				}
				asnNum, err := strconv.ParseUint(record[3], 10, 32)
				if err != nil {
					log.Fatalln(err)
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

	writer, err := mmdbwriter.New(
		mmdbwriter.Options{
			DatabaseType: "GeoLite2-ASN",
			Description:  map[string]string{"en": "IPinfo GeoLite2 ASN database"},
			RecordSize:   24,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
	for asn, info := range asnMap {
		for _, cidr := range info.CIDRs {
			_, ipNet, err := net.ParseCIDR(cidr)
			if err != nil {
				fmt.Println("Error parsing CIDR: ", err)
				continue
			}
			record := mmdbtype.Map{
				"autonomous_system_number":       mmdbtype.Uint32(asn),
				"autonomous_system_organization": mmdbtype.String(info.Name),
			}
			err = writer.Insert(ipNet, record)
			if err != nil {
				fmt.Println("Error inserting record: ", err)
				continue
			}
		}
	}

	filename := "ip2location.asn.mmdb"
	f, err := os.OpenFile(filepath.Join(filename), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = writer.WriteTo(f)
	if err != nil {
		log.Fatalln(err)
	}
}
