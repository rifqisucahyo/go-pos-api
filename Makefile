# Variabel untuk nama binary
BINARY_NAME=redsync

# Daftar file sumber
SOURCES=$(wildcard *.go)

# Aturan default, akan menjalankan program
all: run

# Aturan untuk membuat binary
build:
	go build -o $(BINARY_NAME) $(SOURCES)

# Aturan untuk menjalankan program
run: build
	./$(BINARY_NAME)

# Aturan untuk membersihkan hasil build
clean:
	go clean
	rm -f $(BINARY_NAME)
