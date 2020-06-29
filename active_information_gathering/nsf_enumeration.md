# Network File Share (NFS) Enumeration

NFS is a distributed file system protocol that allows a user on a client to access files over a network as if they were locally mounted. Its usually implemented on Linux, and also usually insecure.

## NFS Scanning

Port mapper and RPC bind run on TCP port 111, RPC bind maps RPC services to ports, and processes notify RPC bind when they start and register their ports and program numbers they expect to serve. RPC bind then contacts that port with its program number, the service is then redirected to the correct port so it can communicate. 

Scan a subnet for systems running RPC bind on port 111
> nmap -sV -p 111 --script=rpcinfo 10.10.10.1-254

## Nmap NFS NSE Scrips

```
[21:27]>>root@kali>>~>>
λ ls -l /usr/share/nmap/scripts/nfs* | cut -d"/" -f6
nfs-ls.nse
nfs-showmount.nse
nfs-statfs.nse
```
Run all three NFS enumeration scripts on a single host
> nmap -p 111 --script=nfs* 10.10.10.155

Mount exposed home directory
> mount -o nolock 10.10.10.155:/home ~/home/

Create user with specific UUID to allow access to restricted mounted files
On local system
> adduser tempuser
>
Change the user's UUID from 1001 to 10014
>sed -i -e 's/1001/1014/g' /etc/passwd
**-i**: Replace file in-place
**-e**: Execute script
**'s/1001/1014/g'** : Replaces all occurrences of 1001 with 1014 inside */etc/passwd *