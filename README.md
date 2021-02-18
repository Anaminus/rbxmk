# rbxmk
`rbxmk` is a command-line tool for manipulating Roblox files.

rbxmk is useful for development workflows that involve many separate files. If
your project is organized into a number of files, such as Lua files for
scripting and model files for assets, rbxmk makes it simple to combine these
files into a final product, be it a game, plugin, model, and so on.

## Installation
**This project is unstable! Use at your own risk!**

rbxmk can also be installed manually.

1. [Install Go](https://golang.org/doc/install)
2. [Install Git](http://git-scm.com/downloads)
3. Using a shell with Git (such as Git Bash), run the following command:

```bash
go install github.com/anaminus/rbxmk/rbxmk@latest
```

If you installed Go correctly, this will install the latest version of rbxmk to
`$GOPATH/bin`, which will allow you run it directly from a shell.

## Usage
The [Documentation page](doc/README.md) provides a complete reference on how
rbxmk is used.
