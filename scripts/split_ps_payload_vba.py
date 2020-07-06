from sys import argv
from sys import stdout

with open(argv[1]) as f:
        shellcode = f.read()
n = 50

for i in range(0, len(shellcode), n):
    stdout.write("Str = Str + " + '"' + shellcode[i:i+n] + '"' + '\n')
