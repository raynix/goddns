package goddns

import (
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
	cf "github.com/cloudflare/cloudflare-go"
)

// Login returns a CloudFlare client pointer
func Login() *cf.API {
	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if hasError(err) {
		log.Fatalf("Can't login!")
	}
	return api
}

// GetZoneID returns the zone ID of a DNS zone
func GetZoneID(api *cf.API, domain string) string {
	zid, err := api.ZoneIDByName(domain)
	if hasError(err) {
		log.Fatalf("Can't get zone id for %v\n", domain)
	}
	return zid
}

// GetRootRecord returns a DNSRecord struct of the root record of the zone
func GetRootRecord(api *cf.API, domain string) cf.DNSRecord {
	zid := GetZoneID(api, domain)
	record := cf.DNSRecord{
		Name:   domain,
		Type:   "A",
		ZoneID: zid,
	}
	recordResults, err := api.DNSRecords(zid, record)
	if len(recordResults) != 1 {
		log.Fatalf("Expected only 1 root record. Error: %v\n", err)
	}
	return recordResults[0]
}
