#include <stdio.h>
#include <stdlib.h>

void die(char* message) {
    fprintf(stderr, message);
    exit(EXIT_FAILURE);
}