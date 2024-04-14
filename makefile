# Makefile

# Define the name of the binary
BINARY_NAME=app

# Define the source files
SOURCES=main.go headers.go readiness.go reset.go

# Define the build command
build:
	go build -o $(BINARY_NAME) $(SOURCES)

# Define the clean command
clean:
	rm -f $(BINARY_NAME)
