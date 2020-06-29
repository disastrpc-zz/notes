# Scanning for the NetBIOS service

## Nmap
Scan entire subnet for SMB and NetBIOS services
> nmap -v -p 139,445 -oG smb.txt 10.10.10.1-255

*NSE SMB Scripts:*
```
[21:12]>>root@kali>>~>>
λ ls -l /usr/share/nmap/scripts/smb* | cut -d"/" -f6
smb2-capabilities.nse
smb2-security-mode.nse
smb2-time.nse
smb2-vuln-uptime.nse
smb-brute.nse
smb-double-pulsar-backdoor.nse
smb-enum-domains.nse
smb-enum-groups.nse
smb-enum-processes.nse
smb-enum-services.nse
smb-enum-sessions.nse
smb-enum-shares.nse
smb-enum-users.nse
smb-flood.nse
smb-ls.nse
smb-mbenum.nse
smb-os-discovery.nse
smb-print-text.nse
smb-protocols.nse
smb-psexec.nse
smb-security-mode.nse
smb-server-stats.nse
smb-system-info.nse
smb-vuln-conficker.nse
smb-vuln-cve2009-3103.nse
smb-vuln-cve-2017-7494.nse
smb-vuln-ms06-025.nse
smb-vuln-ms07-029.nse
smb-vuln-ms08-067.nse
smb-vuln-ms10-054.nse
smb-vuln-ms10-061.nse
smb-vuln-ms17-010.nse
smb-vuln-regsvc-dos.nse
smb-vuln-webexec.nse
smb-webexec-exploit.nse
```

Execute Nmap OS discovery script
> nmap -v -p 139,445 --script=smb-os-discovery 10.10.10.223

## nbtscan 

> nbtscan -r 10.10.10.0/24
-r : Specifies that this scan will be perform on port UDP 137, used to query the NetBIOS name service for valid NetBIOS names. 
