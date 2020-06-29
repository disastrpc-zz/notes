## Active Port Scanning

## TCP nc scan
Perform a port scan on an IP address with a range of ports
> nc -nvv -w 1 -z 10.10.10.11 3388-3390

## Nmap scanning note

Nmap scans will scan the first 1000 ports which are usually the most common. This can cause a lot of traffic so it is important to be aware.


