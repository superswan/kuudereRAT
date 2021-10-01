#include <stdio.h>
#include <stdlib.h>
#include "die.h"

char* get_uuid_machine_id(){
    char* uuid = malloc(36 * sizeof(char));

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

    return uuid;
}