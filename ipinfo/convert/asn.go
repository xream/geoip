package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/maxmind/mmdbwriter"
	"github.com/maxmind/mmdbwriter/mmdbtype"
	"github.com/oschwald/maxminddb-golang"
)

type IPInfo struct {
	ASN    string `maxminddb:"asn"`
	Name   string `maxminddb:"name"`
	Domain string `maxminddb:"domain"`
}

func main() {
	db, err := maxminddb.Open("asn.mmdb")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	type ASNInfo struct {
		Name  string
		CIDRs []string
	}

	asnMap := make(map[uint]ASNInfo)
	networks := db.Networks(maxminddb.SkipAliasedNetworks)

	for networks.Next() {
		var asn IPInfo
		network, err := networks.Network(&asn)
		if err != nil {
			fmt.Println("Error decoding network: ", err)
			continue
		}
		asnNumber, err := strconv.ParseUint(strings.TrimPrefix(asn.ASN, "AS"), 10, 64)
		if err != nil {
			fmt.Println("Error parsing ASN: ", err)
			continue
		}

		info, exists := asnMap[uint(asnNumber)]
		if !exists {
			info = ASNInfo{
				Name:  asn.Name,
				CIDRs: []string{},
			}
		}
		info.CIDRs = append(info.CIDRs, network.String())
		asnMap[uint(asnNumber)] = info
	}

	writer, err := mmdbwriter.New(
		mmdbwriter.Options{
			DatabaseType: "GeoLite2-ASN",
			Description:  map[string]string{"en": "IPinfo GeoLite2 ASN database"},
			RecordSize:   24,
		},
	)
	if err != nil {
		fmt.Println(err)
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

	filename := "new-asn.mmdb"
	f, err := os.OpenFile(filepath.Join(filename), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	_, err = writer.WriteTo(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("New MMDB file created: ", filename)
}
