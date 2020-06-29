# DNS Enum - Interacting with a DNS server

## Basic queries
Provides a hostname > IP resolution.
> host www.megacorp.one

Retrieves mx (or email) records from DNS server.
-t switch allows you to specify record type to use.
> host -t mx megacorpone.com

## DNS Zone Transfers
Attempt to replicate the DNS zone from ns1.megacorpone.com.
> host -l megacorpone.com ns1.megacorpone.com

## Forward lookup brute force
Attempt to resolve DNS host names using ip list
> for ip in $(cat iplist.txt); do host $ip.megacorpone.com; done