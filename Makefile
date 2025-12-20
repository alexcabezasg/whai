# Makefile para compilar y mover el ejecutable 'whai'

install:
	go build -o whai ./cmd/whai
	sudo mv whai /usr/local/bin/