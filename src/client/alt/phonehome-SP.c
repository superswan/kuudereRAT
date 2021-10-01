#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <unistd.h>
#include <arpa/inet.h>

#include "die.h"

int main(int argc, char *argv[])
{
    int s;
    char uuid[36];

    struct sockaddr_in server;

    server.sin_addr.s_addr = inet_addr("173.44.55.179");
    server.sin_family = AF_INET;
    server.sin_port = htons(23855);
    
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

    if ((s=socket(AF_INET, SOCK_STREAM, IPPROTO_TCP)) < 0)
    {
        die("Error creating socket"); 
    }

    if (connect(s, (struct sockaddr *)&server, sizeof(server)) < 0)
    {
        die("connection error\n");
        continue;
    }
    
    if(send(s, uuid, strlen(uuid), 0) < 0)
    {
        die("error while sending transmission\n");
        continue;
    }

    shutdown(s, 2);
    
    return 0;
}

