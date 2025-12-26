# Whai 🫤 ➡️ 🤖 ➡️ 😃

**Whai** (yes, pronounced “_why_”) is your new best friend in the terminal. You run a command, it fails, and **instead of banging your head on the desk**, you just type:

```
whai
```


**whai©️** reads what went wrong and **helps you understand why it failed and how to fix it using AI**. Simple, smart, and right where you need it — your shell. 🚀

## What Whai Actually Does 🎯

Whai is a CLI assistant for debugging failed commands. When a command in bash fails, Whai helps you:

✍️ See a natural‑language explanation of the error

🛠️ Get suggestions for how to fix it

🤔 Understand why the failure occurred

Everything happens inside your terminal — no web chat windows, no alt‑tabbing to a browser.

> [!IMPORTANT]  
> Whai works only with bash. Other shells are not supported.

## Tech Stack ⚙️

- Go — You need Go installed to build from source.

Whai itself is a compiled Go CLI tool, designed to be simple, fast, and terminal‑native.

## Installation & Setup 🛠️
1. Clone the Repo
```
git clone https://github.com/alexcabezasg/whai.git
cd whai
```

2. Build & Install
```
make install bash
```

This compiles Whai and installs the binary in `usr/local/bin`.

3. Initial Setup
```
whai setup
```

There you will need to set up whai needed credentials

### Configuration ⚡

The configuration has the following format:
```json
{
  "only_suggest": false,
  "model": "openai",
  "debug_mode": false,
  "models_configuration": {
    "openai": {
      "url": "",
      "api_key": ""
    }
  }
}
```
- ``only-suggest``: This is used to filter the response of whai to only show the suggestion. No more explanation. Just YOLO mode.
- ``model``: This is the model that whai will use to retrieve the needed information.

> [!IMPORTANT]  
> Whai works only with openai and gpt-5-nano model. Other model families are not supported. Maybe you want to contribute?.

- ``debug_mode``: Its just for debugging issues. It will enable debug logs that contains relevant information about events and data that whai is handling in each request.
- ``models_configuration``: Just a set of credentials

You don't need to change anything, just set the api_key that you should get from openai

## What You Can Do With Whai (Examples) 🚀
### Fix a Failed Command

You run something like:
```shell
cd directoy
bash: cd: directoy: folder doesn't exists
```
Just run ``whai``
```shell
whai

whai
                                                                                                                                                                                                                                                                                                           
Summary: Cannot cd to the target directory due to a typo.                                                                                                                                                                                                                                                  
                                                                                                                                                                                                                                                                                                           
Root Cause: Misspelled or non-existent directory name 'directoy'.                                                                                                                                                                                                                                          
                                                                                                                                                                                                                                                                                                           
Suggestion: cd directory                                           
```

That's it!

## Other available commands
```shell
whai setup [options]
options:
  --only-suggest="true|false"
  --model="openai"
  --debug-mode="true|false"
```
```shell
whai help
```

## Contributing 🤝

Want to improve Whai?

- Fork the repo

- Create a feature branch (git checkout -b my‑feature)

- Make your changes

- Commit with a message (`git commit -m "✨ add magical explanation feature"`)

- Push and open a Pull Request

Contributions are welcome! Just keep it Go‑idiomatic. 😉

## License 📄

Check the repository’s LICENSE file for details.