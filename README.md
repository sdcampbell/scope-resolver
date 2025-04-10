# scope-resolver
Expands a pentest scope by taking an input list which includes IP addresses and parses DNS and TLS SAN hostnames and merges the results with the original scope list. 

Why? Because you need a list of vhosts, hostnames, and subdomains to properly enumerate web applications running behind a load balancer/proxy on an IP address. 

If you only scan and check DNS records, you'll miss applications load balanced on an IP address that don't have DNS records for the application. You may find web applications intended to be protected behind a CDN WAF where their firewall isn't limiting access to CDN IP address ranges and get lucky! I've also found old, forgotten web applications that sat behind load balancers using this tool which the customer didn't realize was exposed to the Internet.

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

This tool works really well together with [nmapurls](https://github.com/sdcampbell/nmapurls). Given an nmap scan xml of IP addresses, some of which do not resolve in DNS, let's find some vhosts via TLS certificates:

```
$ cat nmap.xml | nmapurls
http://83.244.225.132:443
https://83.244.225.136:443
http://83.244.225.138:443
https://83.244.225.154:443
https://83.244.225.161:443

$ cat nmap.xml | nmapurls | scope-resolver
[SSL-SAN] https://83.244.225.136:443 aglctx.avios.com
[SSL-SAN] https://83.244.225.136:443 avbhxdiis01.avios.com
[SSL-SAN] https://83.244.225.136:443 avcbxdiis01.avios.com
[SSL-CN] https://83.244.225.136:443 aglctx.avios.com
[SSL-SAN] https://83.244.225.154:443 gp.avios.com
[SSL-CN] https://83.244.225.154:443 gp.avios.com
[SSL-SAN] https://83.244.225.161:443 *.avios.com
[SSL-SAN] https://83.244.225.161:443 avios.com
[SSL-CN] https://83.244.225.161:443 *.avios.com
[SSL-SAN] http://83.244.225.138:443 adfs.avios.com
[SSL-CN] http://83.244.225.138:443 adfs.avios.com
[SSL-CN] http://83.244.225.132:443 External_Airmiles VPN Certificate
[DNS-PTR] 83.244.225.161 83-244-225-161.cust-83.exponential-e.net
[DNS-PTR] 83.244.225.138 83-244-225-138.cust-83.exponential-e.net
[DNS-PTR] 83.244.225.154 83-244-225-154.cust-83.exponential-e.net
[DNS-PTR] 83.244.225.132 83-244-225-132.cust-83.exponential-e.net
[DNS-PTR] 83.244.225.136 83-244-225-136.cust-83.exponential-e.net
```
