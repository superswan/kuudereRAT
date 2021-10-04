#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <unistd.h>
#include <arpa/inet.h>

void die(char* message);

int main(int argc, char *argv[])
{
    int s, port;
    int retry_count;
    char* host = argv[1];
    port = strtol(argv[2], NULL, 10);
    char uuid[36], server_reply[36];

    struct sockaddr_in server;

    server.sin_addr.s_addr = inet_addr(host);
    server.sin_family = AF_INET;
    server.sin_port = htons(port);
    
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
            ;
        } else {
            printf("%s", server_reply);
        }

        retry_count = 0;
        shutdown(s, 2);
        sleep(33);
    }
    
    return 0;
}

void die(char* message) {
    fprintf(stderr, message);
    exit(EXIT_FAILURE);
}

