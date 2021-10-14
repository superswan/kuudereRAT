#module_input and module_output will be set by enviornment variables in future
module_input = revshell.go
module_output = revshell.so
LDFLAGS = "-X 'main.Host=$(LHOST)'"
PLATFORM = -buildmode=c-shared

build:
	go build -o $(module_output) $(PLATFORM) -ldflags $(LDFLAGS) $(module_input)

module:
	go build -o $(module_output) $(PLATFORM) $(module_input)

clean:
	rm revshell.h revshell.so
