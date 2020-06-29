# Password Cracking

## Wordlist Creation
**cewl**

Web scraper for creation of customized dictionaries using keywords gathered from the website. 
```
cewl -d 2 -m 5 -w docswords.txt https://example.com
```

**Crunch**

Craft and generate patterns and wordlists.

```
crunch <min> max<max> <characterset> -t <pattern> -o <output filename>
```
Pattern for birthday:
```
crunch 10 10 -t @@@@@@0728 -o /root/birthdaywordlist.lst
```

**Identify Hash**

Will output possible and less likely posibilites.

```
hash-identifier
```
## Offline Cracking
**John The Ripper**

John is a versatile and powerful hash cracking tool, allowing for password mutations using predefined wordlists.
```
% john -w:wordlist.txt -rules passwdfile
```
One may view the mangling in action via
```
% john -w:wordlist.txt -rules -stdout | less
```
You can also define your own word-mangling rules. For example, edit john.conf and add a section
```
[List.Rules:LinkedIn]
-: ^[nN]^[iI1!]^[dD]^[eE3]^[kK]^[nN]^[iI1!]^[|lL]
-: $[|lL]$[iI1!]$[nN]$[kK]$[eE3]$[dD]$[iI1!]$[nN]
```
Now is
```
% john -w:wordlist.txt -rules=LinkedIn passwdfile
```

Config file for mutations: 

  ```
  nano /etc/john/john.conf
  ```

**Hashcat**

Password cracking tool that uses a GPU in order to perform faster calculations

ex:
```
hashcat -a 0 -m 400 -a 0 -o wppass2.txt hash rockyou.txt
```

## Online Cracking
_Online password attacks are very noisy therefore they should be used carefully and strategically. Blindly running online attacks can lock accounts and generate large amounts of logs._

**Medusa**

Medusa is a fast and modular login brute forcing tool.
```
medusa -u root -P 500-worst-passwords.txt -h 10.10.10.10 -M ssh
```

**Hydra**

Online cracking tool under active development. 
```
 hydra -P password-file.txt -v 10.11.1.219 snmp
```
Can be used to brute-force SSH:
```
hydra -l root -P password-file.txt 10.11.1.219 ssh
```
Supported services: 
```
adam6500 asterisk cisco cisco-enable cvs firebird ftp[s] http[s]-{head|get|post} http[s]-{get|post}-form http-proxy http-proxy-urlenum icq imap[s] irc ldap2[s] ldap3[-{cram|digest}md5][s] memcached mongodb mssql mysql nntp oracle-listener oracle-sid pcanywhere pcnfs pop3[s] postgres radmin2 rdp redis rexec rlogin rpcap rsh rtsp s7-300 sip smb smtp[s] smtp-enum snmp socks5 ssh sshkey svn teamspeak telnet[s] vmauthd vnc xmpp
```

**Ncrack**

Built by the creators of Nmap, Ncrack is the only tool that is able to brute force RDP efficiently. 

```
ncrack -vv --user admin -P password-file.txt rdp://10.11.1.35
```

Modules:

``` 
SSH, RDP, FTP, Telnet, HTTP(S), POP3(S), IMAP, SMB, VNC, SIP, Redis, PostgreSQL, MySQL, MSSQL, MongoDB, Cassandra, WinRM, OWA
```

## Online crackers:

- Cyberchef @ https://gchq.github.io/CyberChef/
- Hashkiller @ https://hashkiller.co.uk/Cracker
- Crack Station@ https://crackstation.net/


