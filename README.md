# scope-resolver
Expands a pentest scope by taking an input list which includes IP addresses and parses DNS and TLS SAN hostnames and merges the results with the original scope list. Why? Because you need a list of hostnames and subdomains to properly enumerate web applications running on an IP address.

Build: `go build scope-resolver.go` or `go run scope-resolver.go [args]`

Usage:

```
Usage of scope-resolver:
  (no arguments)         Read from stdin
  <file_path>            Read from the specified file
  -h or --help           Print this help message

Examples:
  cat ips.txt | scope-resolver
  ./scope-resolver ips.txt
```
