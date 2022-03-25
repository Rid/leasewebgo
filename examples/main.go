package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"leasewebgo"
)

type ErrorMessage struct {
	CorrelationId string `json:"correlationId"`
	ErrorCode     string `json:"errorCode"`
	ErrorMessage  string `json:"errorMessage"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	hostname := "node1.23-106-61-231.example.com"
	fmt.Println("hostname:", hostname)

	localIp := strings.Join(strings.Split(strings.Split(hostname, ".")[1], "-"), ".")
	fmt.Println("localIp:", localIp)

	floatingIp := "173.234.16.25"

	c, err := leasewebgo.NewClient()
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Find Range with matching floating IP
	limit := "50"
	ranges, _, _, err := c.FloatingIps.ListRanges(&leasewebgo.ListOptions{QueryParams: map[string]string{"limit": limit}})
	if err != nil {
		log.Fatal("Error:", err)
	}

	totalCount := ranges.Meta.Totalcount
	offset := 0
	ipRangeId := ""

out:
	for totalCount > 0 {
		for _, ipRange := range ranges.Ranges {
			log.Println("ipRange:", ipRange.Id)
			// convert string to IPNet struct
			_, ipv4Net, err := net.ParseCIDR(ipRange.Range)
			if err != nil {
				log.Fatal(err)
			}

			// convert IPNet struct mask and address to uint32
			// network is BigEndian
			mask := binary.BigEndian.Uint32(ipv4Net.Mask)
			start := binary.BigEndian.Uint32(ipv4Net.IP)

			// find the final address
			finish := (start & mask) | (mask ^ 0xffffffff)

			// loop through addresses as uint32
			for i := start; i <= finish; i++ {
				// convert back to net.IP
				ip := make(net.IP, 4)
				binary.BigEndian.PutUint32(ip, i)
				if ip.String() == floatingIp {
					log.Println("IP matches:", ip)
					ipRangeId = ipRange.Id
					break out
				}
			}
			totalCount -= 1
			offset += 1
		}
		ranges, _, _, err = c.FloatingIps.ListRanges(&leasewebgo.ListOptions{QueryParams: map[string]string{"limit": limit, "offset": strconv.Itoa(offset)}})
	}

	_, _, body, err := c.FloatingIps.UpdateRange(ipRangeId, strings.Join([]string{floatingIp, "32"}, "_"), &leasewebgo.RangeUpdateRequest{AnchorIp: localIp}, nil)
	if err != nil {
		if len(body) > 0 {
			errorMessage := &ErrorMessage{}
			json.Unmarshal(body, errorMessage)
			log.Println(errorMessage.ErrorMessage)
			if errorMessage.ErrorCode == "404" {
				if strings.Contains(errorMessage.ErrorMessage, "does not exist") {
					log.Println("Creating new range")
					_, _, _, err = c.FloatingIps.CreateRange(ipRangeId, &leasewebgo.RangeCreateRequest{FloatingIp: floatingIp, AnchorIp: localIp}, nil)
					if err != nil {
						log.Println("Error:", err)
					}
				}
			}
		}
	} else {
		log.Println("Updated floating IP successfully")
	}
}
