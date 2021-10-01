#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <netinet/in.h>
#include <netdb.h>
#include <unistd.h>
#include <arpa/inet.h>


char* resolveIP(char *hostname) {
    struct addrinfo hints, *res, *p;
    int status;
    char ipstr[INET6_ADDRSTRLEN];
    char *ipaddr = malloc(INET6_ADDRSTRLEN);

   memset(&hints, 0, sizeof hints);
   hints.ai_family = AF_UNSPEC;
   hints.ai_socktype = SOCK_STREAM;

    if ((status = getaddrinfo(hostname, NULL, &hints, &res)) != 0) {
        return NULL;
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