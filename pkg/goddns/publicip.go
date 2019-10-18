package goddns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func hasError(err error) bool {
	if err != nil {
		fmt.Printf("%v\n", err)
		return true
	}
	return false
}

func httpGetBody(url string) (body []byte) {
	resp, err := http.Get(url)
	if hasError(err) {
		return []byte{}
	}
	defer resp.Body.Close()
	bodyBuffer, err := ioutil.ReadAll(resp.Body)
	if hasError(err) {
		return []byte{}
	}
	return bodyBuffer
}

type publicIPSniffer interface {
	publicIP() string
}

// HttpBinOrg represents httpbin.org
type HttpBinOrg struct {
	// {
	//   "origin": "49.183.63.37, 49.183.63.37"
	// }
	Origin string `json:"origin"`
}

func (s HttpBinOrg) publicIP() string {
	url := "https://httpbin.org/ip"
	json.Unmarshal(httpGetBody(url), &s)
	return strings.Split(s.Origin, ",")[0]
}

// APIipify represents api.ipify.org
type APIipify struct {
	// {"ip":"101.116.192.67"}
	IP string `json:"ip"`
}

func (s APIipify) publicIP() string {
	url := "https://api.ipify.org/?format=json"
	json.Unmarshal(httpGetBody(url), &s)
	return s.IP
}

type dummySniffer struct {
	IP string `json:"ip"`
}

func (s dummySniffer) publicIP() string {
	dummy := `{"ip":"101.116.192.67"}`
	json.Unmarshal([]byte(dummy), &s)
	return s.IP
}

// GetPublicIP returns current public IP using some service provider
func GetPublicIP(sniffer publicIPSniffer) string {
	return sniffer.publicIP()
}
