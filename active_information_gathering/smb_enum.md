## SMB/RPC Enumeration

**Nmap smb-enum script**

Collection of scripts to enumerate public SMB shares, users and services.

Inludes smb-enum-shares, smb-enum-users and smb-enum-services.

```
nmap -sV --script smb-enum-*
```
**enum4linux**

Discover Windows and Samba servers on specified subnet. Can also discover netbios names and client workgroup/domain.
```
enum4linux -a 10.1.62.7
```