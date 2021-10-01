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
    int retry_count;
    char uuid[36], server_reply[36];

    struct sockaddr_in server;

    server.sin_addr.s_addr = inet_addr("10.1.0.199");
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

        if(recv(s, server_reply, sizeof(server_reply), 0) < 0)
        {
            printf("error receiving message or no tasks");
        } else {
            printf("%s", server_reply);
        }

        retry_count = 0;
        shutdown(s, 2);
        sleep(33);
    }
    
    return 0;
}

