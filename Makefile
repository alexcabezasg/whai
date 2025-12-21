
install:
	go build -o whai ./cmd/whai
	sudo mv whai /usr/local/bin/
	sudo mkdir ~/.whai
	sudo touch ~/.whai/config.json
	sudo chmod 666 ~/.whai/config.json
	sudo cat config-template.json >> ~/.whai/config.json