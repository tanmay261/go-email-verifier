package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, DMARCrecord \n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error :Cannot read from input: %v\n", err)
	}
}

func checkDomain(domain string) {

	var hasDMARC, hasMX, hasSPF bool
	var spfRecord, DMARCrecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error:  %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error:  %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error:  %v\n", err)
	}

	for _, records := range dmarcRecords {
		if strings.HasPrefix(records, "v=DMARC1") {
			hasSPF = true
			DMARCrecord = records
			break
		}
	}

	fmt.Printf("domain:- %v\nhasMX:- %v\nhasSPF:- %v\nspfRecord:- %v\nhasDMARC:- %v\nDMARCrecord:- %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, DMARCrecord)
}
