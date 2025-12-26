install bash:
	@echo "Installing whai..."
	@go build -o whai ./cmd/whai
	@sudo mv whai /usr/local/bin/
	@mkdir -p ~/.whai
	@echo "Adding custom whai logic to save last failed command..."
	@if [ ! -f ~/.bashrc.bak ]; then cp ~/.bashrc ~/.bashrc.bak; fi
	@if ! grep -q "LAST_FAILED_CMD" ~/.bashrc; then \
		echo "" >> ~/.bashrc; \
		echo "# ==============WHAI SAVE LAST FAILED COMMAND LOGIC ===============" >> ~/.bashrc; \
		echo "PROMPT_COMMAND=\"last_status=\$$?; if [ \$$last_status -ne 0 ]; then LAST_FAILED_CMD=\$$(history 1 | sed 's/^[ ]*[0-9]\+[ ]*//'); export LAST_FAILED_CMD; fi\"" >> ~/.bashrc; \
		echo "# =============================" >> ~/.bashrc; \
	fi
	@bash -c "source ~/.bashrc"
	@echo "Installation complete."