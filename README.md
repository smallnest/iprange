package `iprange` provides some methods to check whether a [net.IP](https://golang.org/pkg/net/#IP) is contained in some IP ranges.

The ip range is defined via [CIRD](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing), for example, `216.249.16.0/20` and `2a01:5a80::/32`.
There are different methods for IPv4 or IPv6 processing.

If you want to use this library, you first get it:
```sh
go get github.com/smallnest/iprange
```

Then you can parse a IPv4 or IPv6 range list (assume the list is sorted, otherwise you must sort it like search_test.go):
```go
ipRanges := ParseIPV4RangeFromFile("cidr_ipv4_test.data")
```

Now you can check a whether net.IP is contained in this list:
```go
existed := IPv4Contains(ipRanges, ip)
```

It uses [binary search algorithm](https://en.wikipedia.org/wiki/Binary_search_algorithm) to check so the check is very effective method.

Binary search runs in at worst logarithmic time, making O(log n) comparisons, where n is the number of elements in the list and log is the logarithm. Binary search takes only constant (O(1)) space, meaning that the space taken by the algorithm is the same for any number of elements in the array.

- Worst-case performance:       `O(log n)`
- Best-case performance:        `O(1)`
- Average performance:	        `O(log n)`
- Worst-case space complexity:  `O(1)`

`bit trie` algorithm should be more effective but I have not implemented it.

I found other implementations ~~so you prefer those implementation with radix tree~~:
- [go-net-radix](https://github.com/thekvs/go-net-radix)
- [nradix](https://github.com/asergeyev/nradix)
- [Longest Prefix Match algorithm in Go part 1](https://fredhsu.wordpress.com/2014/06/09/longest-prefix-match-algorithm-in-go-part-1/)
- [iptree](https://github.com/zmap/go-iptree)

The below is benchmark of [iprange](https://github.com/smallnest/iprange), [go-net-radix](https://github.com/thekvs/go-net-radix), [nradix](https://github.com/asergeyev/nradix):

[nradix](https://github.com/asergeyev/nradix) and [iprange](https://github.com/smallnest/iprange) are much better than [go-net-radix](https://github.com/thekvs/go-net-radix).


```
BenchmarkIPv4Contains-8          	  500000	      3387 ns/op	       0 B/op	       0 allocs/op
BenchmarkIPv4Contains_Radix-8    	    5000	    223566 ns/op	     696 B/op	      70 allocs/op
BenchmarkIPv4Contains_NRadix-8   	  500000	      3439 ns/op	     192 B/op	      34 allocs/op

BenchmarkIPv6Contains-8          	  300000	      5835 ns/op	     720 B/op	      18 allocs/op
BenchmarkIPv6Contains_Radix-8    	    5000	    224394 ns/op	     728 B/op	      45 allocs/op
BenchmarkIPv6Contains_NRadix-8   	  300000	      4315 ns/op	     800 B/op	      33 allocs/op
```
