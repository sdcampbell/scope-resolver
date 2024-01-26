# scope-resolver
Expands a pentest scope by taking an input list which includes IP addresses and parses DNS and TLS SAN hostnames and merges the results with the original scope list. Why? Because you need a list of hostnames and subdomains to properly enumerate web applications running on an IP address.

Install: `go install github.com/sdcampbell/scope-resolver@latest`

Usage:

```
Usage of scope-resolver:
  -p int
    	Port to bother the specified DNS resolver on (default 53)
  -protocol string
    	Protocol for DNS lookups (tcp or udp) (default "udp")
  -r string
    	IP of DNS resolver for lookups
  -t int
    	numbers of threads (default 32)

Examples:
  cat ips.txt | scope-resolver
  ./scope-resolver ips.txt
```
