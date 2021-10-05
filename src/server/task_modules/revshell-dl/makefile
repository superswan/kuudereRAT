build: library 
	gcc -o shell shell.c -ldl

library:
	go build -o revshell.so -buildmode=c-shared revshell.go

clean:
	rm shell revshell.so revshell.h