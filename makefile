# Makefile

# Define the name of the binary
BINARY_NAME=out 
DB=database.json

# Define the source files
SOURCES=*.go

# Define the build command
build:
	make clean
	go build -o $(BINARY_NAME) $(SOURCES)
	./out
	

# Define the clean command
clean:
	rm -f $(BINARY_NAME)
	rm -f $(DB)
	

