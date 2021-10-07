#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <unistd.h>
#include <arpa/inet.h>

void die(char *error);
char* resolveIP(char *hostname);
int main(int argc, char *argv[])
{
    int s;
    int retry_count;
    char uuid[36];

    struct sockaddr_in server;
    char* srvaddr = resolveIP("hostname");

    server.sin_addr.s_addr = inet_addr(srvaddr);
    server.sin_family = AF_INET;
    server.sin_port = htons(23855);

    free(srvaddr);
        
    FILE *fp=fopen("/var/lib/dbus/machine-id", "r");
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

    retry_count = 0;
    while(1){
        if ((s=socket(AF_INET, SOCK_STREAM, IPPROTO_TCP)) < 0)
        {
           printf("Error creating socket"); 
           retry_count += 1;
           sleep(30 * retry_count);
        }

        if (connect(s, (struct sockaddr *)&server, sizeof(server)) < 0)
        {
            printf("connection error\n");
            retry_count += 1;
            sleep(30 * retry_count);
            continue;
        }
        
        if(send(s, uuid, strlen(uuid), 0) < 0)
        {
            printf("error while sending transmission\n");
            retry_count += 1;
            sleep(30 * retry_count);
            continue;
        }

        retry_count = 0;
        shutdown(s, 2);
        sleep(33);
    }
    
    return 0;
}

void die(char *error) {
    fprintf(stderr, "%s\n", error);
    exit(1);
}

char* resolveIP(char *hostname) {
    struct addrinfo hints, *res, *p;
    int status;
    char ipstr[INET6_ADDRSTRLEN];
    char *ipaddr = malloc(INET6_ADDRSTRLEN);

   memset(&hints, 0, sizeof hints);
   hints.ai_family = AF_UNSPEC;
   hints.ai_socktype = SOCK_STREAM;

    if ((status = getaddrinfo(hostname, NULL, &hints, &res)) != 0) {
        die("Can't resolve hostname");
    }               

    for(p = res; p != NULL; p = p->ai_next) {
        void *addr;
        char *ipver;

        if (p->ai_family==AF_INET) {
            struct sockaddr_in *ipv4 = (struct sockaddr_in *)p->ai_addr;
            addr = &(ipv4->sin_addr);
            ipver = "IPv4";
        }
        
        inet_ntop(p->ai_family, addr, ipstr, sizeof(ipstr));
    }
    
    snprintf(ipaddr, INET6_ADDRSTRLEN, "%s", ipstr);
    freeaddrinfo(res); 

    
    return ipaddr;
}