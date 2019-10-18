package goddns

import (
  "testing"
)

func TestHttpBinOrg(t *testing.T) {
  s := HttpBinOrg{}
  ip := GetPublicIP(s)
  if len(ip) == 0 {
    t.Error("Got empty public IP from httpbin.org")
  }
  t.Logf("Got public IP: %v\n", ip)
}

func TestIpify(t *testing.T) {
  s := APIipify{}
  ip := GetPublicIP(s)
  if len(ip) == 0 {
    t.Error("Got empty public IP from httpbin.org")
  }
  t.Logf("Got public IP: %v\n", ip)
}

func TestPublicIPSniffer(t *testing.T) {
  dummy := dummySniffer{}
  ip := GetPublicIP(dummy)
  if len(ip) == 0 {
    t.Error("Dummy Sniffer test failed")
  }
  t.Logf("Got public IP: %v\n", ip)
}
