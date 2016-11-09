package iprange

import "net"

// IPv4Contains is used to check whether the IPv4 is in the IPRanges
func IPv4Contains(ipRanges []*IPV4Range, ip net.IP) bool {
	if ipRanges == nil || len(ipRanges) == 0 {
		return false
	}
	maxIndex := len(ipRanges) - 1
	if maxIndex == 0 {
		return ipRanges[0].IPNet.Contains(ip)
	}

	ipNum := ipv4toInt(ip)

	low, high := 0, maxIndex
	for low <= high {
		mid := low + (high-low)/2
		if ipRanges[mid].IPNet.Contains(ip) {
			return true
		}

		if ipRanges[mid].Start < ipNum {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == 0 {
		return ipRanges[0].IPNet.Contains(ip)
	}
	if low == maxIndex+1 {
		return ipRanges[maxIndex].IPNet.Contains(ip)
	}

	return ipRanges[low-1].IPNet.Contains(ip)
}

// IPv6Contains is used to check whether the IPv6 is in the IPRanges
func IPv6Contains(ipRanges []*IPV6Range, ip net.IP) bool {
	if ipRanges == nil || len(ipRanges) == 0 {
		return false
	}
	maxIndex := len(ipRanges) - 1
	if maxIndex == 0 {
		return ipRanges[0].IPNet.Contains(ip)
	}

	ipNum := ipv6toInt(ip)

	low, high := 0, maxIndex
	for low <= high {
		mid := low + (high-low)/2
		if ipRanges[mid].IPNet.Contains(ip) {
			return true
		}

		if ipRanges[mid].Start.Cmp(ipNum) < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == 0 {
		return ipRanges[0].IPNet.Contains(ip)
	}
	if low == maxIndex+1 {
		return ipRanges[maxIndex].IPNet.Contains(ip)
	}

	return ipRanges[low-1].IPNet.Contains(ip)
}
