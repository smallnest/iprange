package iprange

import (
	"net"
	"testing"

	"github.com/bradfitz/slice"
)

func TestIPv4Contains(t *testing.T) {
	ipRanges := ParseIPV4RangeFromFile("cidr_ipv4_test.data")

	//sort. Go 1.8 supports sort.Slice(things, func(i, j int) bool) but it will be released next year
	slice.Sort(ipRanges, func(i, j int) bool {
		return ipRanges[i].Start < ipRanges[j].Start
	})

	type args struct {
		ipRanges []*IPV4Range
		ip       net.IP
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "MidStart",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("103.67.32.0")},
			want: true,
		},
		{
			name: "MidMid",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("103.67.32.1")},
			want: true,
		},
		{
			name: "MidMissed",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("103.67.100.77")},
			want: false,
		},
		{
			name: "Lowbound",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("3.0.0.0")},
			want: true,
		},
		{
			name: "Upperbound",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("216.255.255.255")},
			want: true,
		},
		{
			name: "Lowbound-1",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("2.255.255.255")},
			want: false,
		},
		{
			name: "Upperbound+1",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("217.0.0.0")},
			want: false,
		},
		{
			name: "First",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("0.0.0.0")},
			want: false,
		},
		{
			name: "Last",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("255.255.255.255")},
			want: false,
		},
	}

	for _, tt := range tests {
		if got := IPv4Contains(tt.args.ipRanges, tt.args.ip); got != tt.want {
			t.Errorf("%q. Contains() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestIPv6Contains(t *testing.T) {
	ipRanges := ParseIPV6RangeFromFile("cidr_ipv6_test.data")

	//sort. Go 1.8 supports sort.Slice(things, func(i, j int) bool) but it will be released next year
	slice.Sort(ipRanges, func(i, j int) bool {
		return ipRanges[i].Start.Cmp(ipRanges[j].Start) < 0
	})

	type args struct {
		ipRanges []*IPV6Range
		ip       net.IP
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "MidStart",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("2607:d200::")},
			want: true,
		},
		{
			name: "MidMid",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("2607:d200::1")},
			want: true,
		},
		{
			name: "MidMissed",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("2607:d201::ffff")},
			want: false,
		},
		{
			name: "Lowbound",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("2001:1800::")},
			want: true,
		},
		{
			name: "Upperbound",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("2a03:cd00:ffff:ffff:ffff:ffff:ffff:ffff")},
			want: true,
		},
		{
			name: "Lowbound-1",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("2001:17ff:ffff:ffff:ffff:ffff:ffff:ffff")},
			want: false,
		},
		{
			name: "Upperbound+1",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("2a03:cd01::")},
			want: false,
		},
		{
			name: "First",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("::")},
			want: false,
		},
		{
			name: "Last",
			args: args{ipRanges: ipRanges, ip: net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")},
			want: false,
		},
	}

	for _, tt := range tests {
		if got := IPv6Contains(tt.args.ipRanges, tt.args.ip); got != tt.want {
			t.Errorf("%q. Contains() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
