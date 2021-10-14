/*
Uses /var/lib/dbus/machine-id as UUID and /etc/machine-id as a failover. 
From Man page: 
It's intended as an opaque, non-human-meaningful, persistent unique identifier 
for a machine (or more precisely an OS installation), used as a lookup key in 
state/configuration storage in the same sorts of places you might be tempted to
use a hostname.
*/

#include <stdio.h>
#include <stdlib.h>
#include "die.h"

char* get_uuid_machine_id(){
    char* uuid = malloc(36 * sizeof(char));

    FILE *fp=fopen("/var/lib/dbus/machine-id", "r");
        if (fp == NULL)
        {
            fclose(fp);
            FILE *fp=fopen("/etc/machine-id", "r");

            if (fp == NULL) 
            {
                die("Error getting machine ID");
            }
        } 
        
        if (fgets(uuid, 33, fp) == NULL)
        {
            die("Error getting machine ID");
        }
        fclose(fp);

    return uuid;
}