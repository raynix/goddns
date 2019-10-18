package main

import (
	"time"
	"fmt"
	"os"
	goddns "github.com/raynix/goddns/pkg/goddns"
)

func main() {
	domains := os.Args[1:]
	client := goddns.Login()
	for ;; {
		time.Sleep(30 * time.Second)
		publicIP := goddns.GetPublicIP(goddns.APIipify{})
		for _, domain := range domains{
			rootRecord := goddns.GetRootRecord(client, domain)
			if publicIP != rootRecord.Content {
				fmt.Printf("Public IP %v is not same as DNS IP %v, will update\n", publicIP, rootRecord.Content)
				rootRecord.Content = publicIP
				err := client.UpdateDNSRecord(rootRecord.ZoneID, rootRecord.ID, rootRecord)
				if err != nil {
					fmt.Printf("%v\n", err)
				}
			}
		}
	}
}
