
install:
	go build -o whai ./cmd/whai
	sudo mv whai /usr/local/bin/
	mkdir ~/.whai