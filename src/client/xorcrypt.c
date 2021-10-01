#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#include "die.h"

char* genkey(){
    char* key[128];
    return key;
}

char* sxorcrypt(char* message, char* key){
    int i, j, c, message_length, key_length, max_size;
    char xor_c;

    message_length = (strlen(key)) - 1;
    key_length = (strlen(key)) - 1;

    printf("%d", message_length); 
    // rotates over the key from LTR to obfuscate XOR
    for (i = 0, j = 0, max_size=500000; i < max_size && (c = getc(message)) != EOF; i++) {
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
        message[i] = xor_c;
    }
    return message; 
} 

int fxorcrypt(char* filename, char* filename_enc, char* key) {
    FILE *fp;
    FILE *fp_enc;
    
    int i, j, c, key_length, max_size;
    char xor_c;

    key_length = (strlen(key)) - 1;
    
    fp = fopen(filename, "r");
    if (fp == NULL) {
        die("Error opening file\n");
    }

    fp_enc = fopen(filename_enc, "w");
    if (fp == NULL) {
        die("Error creating file\n");
    }

    for (i = 0, j = 0, max_size=500000; i < max_size && (c = getc(fp)) != EOF; i++) {
        if (j > key_length) {
            j = 0;
        }

        if (c == 0) {
            xor_c = c;
        } 
        else xor_c = c ^ key[j];

        if (xor_c == 0) {
            xor_c = c;
        }

        j++;
        fputc(xor_c, fp_enc); 
    } 
    
    fclose(fp);
    fclose(fp_enc);

    return 0;
}


