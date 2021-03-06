package main

import (
	"time"
	"log"
	"os"
	"strings"
	"strconv"
	goddns "github.com/raynix/goddns/pkg/goddns"
)

func main() {
	log.SetOutput(os.Stdout)
	domains := os.Args[1:]
	domainsEnv := os.Getenv("GODDNS_DOMAINS")
	dryRun := os.Getenv("DRYRUN") == "true"
	intervalEnv, err := strconv.Atoi(os.Getenv("INTERVAL"))
	interval := 300
	if err == nil {
		interval = intervalEnv
	}
	if len(domainsEnv) > 0 {
		domains = strings.Split(domainsEnv, ",")
	}
	client := goddns.Login()
	for ;; {
		log.Printf("Waiting for %v seconds...", interval)
		time.Sleep(time.Duration(interval) * time.Second)
		publicIP := goddns.GetPublicIP(goddns.APIipify{})
		log.Printf("Current public IP is: %v\n", publicIP)
		for _, domain := range domains{
			log.Printf("Checking zone %v\n", domain)
			rootRecord := goddns.GetRootRecord(client, domain)
			if publicIP != rootRecord.Content {
				log.Printf("Public IP %v is not same as DNS IP %v, will update\n", publicIP, rootRecord.Content)
				rootRecord.Content = publicIP
				if dryRun {
					log.Println("Not updating because DRYRUN is on.")
				} else {
					log.Printf("Updating domain %v now.\n", domain)
					err := client.UpdateDNSRecord(rootRecord.ZoneID, rootRecord.ID, rootRecord)
					if err != nil {
						log.Printf("%v\n", err)
					}
				}
			}
		}
	}
}
