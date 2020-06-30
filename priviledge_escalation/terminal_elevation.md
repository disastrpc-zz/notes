# Terminal Elevation

## If python is installed:
```
echo "import pty; pty.spawn('/bin/bash')" > /tmp/asdf.py

python /tmp/asdf.py
```

## Upgrading from a nc shell using magic (thank you Phineas Fisher)

From a nc reverse shell one can upgrade to a full TTY shell by first using Python's pty module to spawn a PTY shell. Afterwards background the process using Ctrl + Z and list the *$TERM* var as well as all the terminal's settings. All we would need is the xterm and row/columns of the current shell.
```
www-data@Raven:/var/www/html$ python -c 'import pty; pty.spawn("/bin/bash")'
python -c 'import pty; pty.spawn("/bin/bash")'
www-data@Raven:/var/www/html$ ^Z
[2]+  Stopped                 nc -nlvp 4040
root@palantir:~# echo $TERM
xterm-256color
root@palantir:~# stty -a
speed 38400 baud; rows 24; columns 144; line = 0;
intr = ^C; quit = ^\; erase = ^?; kill = ^U; eof = ^D; eol = <undef>; eol2 = <undef>; swtch = <undef>; start = ^Q; stop = ^S; susp = ^Z;
rprnt = ^R; werase = ^W; lnext = ^V; discard = ^O; min = 1; time = 0;
-parenb -parodd -cmspar cs8 -hupcl -cstopb cread -clocal -crtscts
-ignbrk brkint ignpar -parmrk -inpck -istrip -inlcr -igncr icrnl ixon -ixoff -iuclc -ixany imaxbel iutf8
opost -olcuc -ocrnl onlcr -onocr -onlret -ofill -ofdel nl0 cr0 tab0 bs0 vt0 ff0
isig icanon iexten echo echoe echok -echonl -noflsh -xcase -tostop -echoprt echoctl echoke -flusho -extproc
```
Then tell stty to echo all characters to the screen in a raw format. *fg* should bring the background nc session back up and can be restored with *reset*.
```
root@palantir:~# stty raw -echo
root@palantir:~# nc -nlvp 4040
                              reset
```
And set the variables to your own shell's.
```
www-data@Raven:/var/www/html$ export SHELL=bash
www-data@Raven:/var/www/html$ export TERM=xterm-256color
www-data@Raven:/var/www/html$ stty rows 24 columns 144
```
The resulting shell should be a full TTY shell, able to do command completion, interactive sessions and everything else.

