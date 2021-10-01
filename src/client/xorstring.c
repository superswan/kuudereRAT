#include <stdlib.h>
#include <stdio.h>
#include <string.h>

char* sxorcrypt(char* message, char* key){
    int i, j, c, message_length, key_length, max_size;
    char xor_c;

    message_length = (strlen(message)) - 1;
    key_length = (strlen(key)) - 1;
    
    char message_enc[message_length];
    // rotates over the key from LTR to obfuscate XOR
    for (i = 0, j = 0; (i < message_length) && ((c = getchar()) != '\0'); i++) {
        if (j > key_length) {
            j = 0;
        }

        if (c == 0) {
            xor_c = c;
        } else {
            xor_c = c ^ key[j];
        }

        if (xor_c == 0) {
            xor_c = c;
        }

        j++;
        message_enc[i] = xor_c;
    }
    return message_enc; 
} 

int main(int argc, char* argv[]) {
    char* message = argv[1];
    char* key = argv[2];

    char* message_enc = sxorcrypt(message, key);
    printf("%s\t%s\n%s\n",message, key, message_enc);

    return 0;
}