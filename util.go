package iprange

import (
	"encoding/binary"
	"math/big"
	"net"
)

func ipv4toInt(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func inttoIPv4(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}

func ipv6toInt(ipv6Address net.IP) *big.Int {
	ipv6Int := big.NewInt(0)
	ipv6Int.SetBytes(ipv6Address.To16())
	return ipv6Int
}

func inttoIPv6(n *big.Int) net.IP {
	return net.IP(n.Bytes())
}
