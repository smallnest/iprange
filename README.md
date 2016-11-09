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
