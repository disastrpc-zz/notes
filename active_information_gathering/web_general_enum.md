# **Web Enumeration**
## **General Tools**

## Apache
Apache version can be obtained by nmap scans using -sV or http-enum script

## apache-users Perl script

_Install with apt-get install apache-users_

This Perl script will enumerate the usernames on any system that uses Apache with the UserDir module.

Example:
```
root@kali:~# apache-users -h 192.168.1.202 -l /usr/share/wordlists/metasploit/unix_users.txt -p 80 -s 0 -e 403 -t 10
```
_Where -h is host, -l is wordlist, -s disabled ssl, -e specifying error code 403 and -t 10 using 10 threads._

## **Wordpress**

## WPScan

Wordpress security scanner
```
wpscan --url <url> -e ap,at
```


## Nikto

Open source web server scanner providing an overview of interesting attack vectors

```
nikto -Display 1234EP -o report.html -Format htm -Tuning 123bde -host 192.168.0.102
```

