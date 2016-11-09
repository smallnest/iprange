package iprange

import (
	"bufio"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
)

// IPV4Range contains two fields.
// One is the net.IPNet and the the other is a uint32 that is the first IP in this IP Range.
type IPV4Range struct {
	Start uint32
	IPNet *net.IPNet
}

// IPV6Range contains two fields.
// One is the net.IPNet and the the other is a *big.Int that is the first IP in this IP Range.
type IPV6Range struct {
	Start *big.Int
	IPNet *net.IPNet
}

// ParseIPV4RangeFromFile parses a file and returns a slice of IPV4Range.
func ParseIPV4RangeFromFile(fileName string) []*IPV4Range {
	var ipranges []*IPV4Range

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, ipnet, err := net.ParseCIDR(strings.Trim(scanner.Text(), " "))
		if err == nil {
			ipranges = append(ipranges, &IPV4Range{Start: ipv4toInt(ipnet.IP), IPNet: ipnet})
		}
	}

	return ipranges
}

// ParseIPV4Range parses a string and returns a slice of IPV4Range.
func ParseIPV4Range(list string) []*IPV4Range {
	var ipranges []*IPV4Range

	scanner := bufio.NewScanner(strings.NewReader(list))
	for scanner.Scan() {
		_, ipnet, err := net.ParseCIDR(strings.Trim(scanner.Text(), " "))
		if err == nil {
			ipranges = append(ipranges, &IPV4Range{Start: ipv4toInt(ipnet.IP), IPNet: ipnet})
		}
	}

	return ipranges
}

// ParseIPV6RangeFromFile parses a file and returns a slice of IPV6Range.
func ParseIPV6RangeFromFile(fileName string) []*IPV6Range {
	var ipranges []*IPV6Range

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, ipnet, err := net.ParseCIDR(strings.Trim(scanner.Text(), " "))
		if err == nil {
			ipranges = append(ipranges, &IPV6Range{Start: ipv6toInt(ipnet.IP), IPNet: ipnet})
		}
	}

	return ipranges
}

// ParseIPV6Range parses a string and returns a slice of IPV6Range.
func ParseIPV6Range(list string) []*IPV6Range {
	var ipranges []*IPV6Range

	scanner := bufio.NewScanner(strings.NewReader(list))
	for scanner.Scan() {
		_, ipnet, err := net.ParseCIDR(strings.Trim(scanner.Text(), " "))
		if err == nil {
			ipranges = append(ipranges, &IPV6Range{Start: ipv6toInt(ipnet.IP), IPNet: ipnet})
		}
	}

	return ipranges
}
