package iprange

import (
	"io/ioutil"
	"testing"
)

func TestParseIPV4RangeFromFile(t *testing.T) {
	ipRanges := ParseIPV4RangeFromFile("cidr_ipv4_test.data")
	testIPv4(t, ipRanges)
}

func TestParseIPV4Range(t *testing.T) {
	b, err := ioutil.ReadFile("cidr_ipv4_test.data")
	if err != nil {
		t.Error(err)
	}
	ipRanges := ParseIPV4Range(string(b))
	testIPv4(t, ipRanges)
}

func testIPv4(t *testing.T, ipRanges []*IPV4Range) {
	l := len(ipRanges)
	if l != 16048 {
		t.Errorf("size of IP: %d but expected: 16048", l)
	}

	if ipRanges[0].IPNet.IP.String() != "3.0.0.0" {
		t.Errorf("First IP: %s but expected: 3.0.0.0", ipRanges[0].IPNet.IP.String())
	}

	if ipRanges[l-1].IPNet.IP.String() != "216.255.240.0" {
		t.Errorf("First IP: %s but expected: 216.255.240.0", ipRanges[l-1].IPNet.IP.String())
	}
}

func TestParseIPV6RangeFromFile(t *testing.T) {
	ipRanges := ParseIPV6RangeFromFile("cidr_ipv6_test.data")
	testIPv6(t, ipRanges)
}

func TestParseIPV6Range(t *testing.T) {
	b, err := ioutil.ReadFile("cidr_ipv6_test.data")
	if err != nil {
		t.Error(err)
	}
	ipRanges := ParseIPV6Range(string(b))
	testIPv6(t, ipRanges)
}

func testIPv6(t *testing.T, ipRanges []*IPV6Range) {
	l := len(ipRanges)
	if l != 2918 {
		t.Errorf("size of IP: %d but expected: 2918", l)
	}

	if ipRanges[0].IPNet.IP.String() != "2001:1800::" {
		t.Errorf("First IP: %s but expected: 2001:1800::", ipRanges[0].IPNet.IP.String())
	}

	if ipRanges[l-1].IPNet.IP.String() != "2a03:cd00::" {
		t.Errorf("First IP: %s but expected: 2a03:cd00::", ipRanges[l-1].IPNet.IP.String())
	}
}
