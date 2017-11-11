# rbxmk

`rbxmk` is a command-line tool for manipulating Roblox files.

rbxmk is useful for development workflows that involve many separate files. If
your project is organized into a number of files, such as Lua files for
scripting and model files for assets, rbxmk makes it simple to combine these
files into a final product, be it a game, plugin, model, and so on.

## Installation

**This project is unstable! Use at your own risk!**

1. [Install Go](https://golang.org/doc/install)
2. [Install Git](http://git-scm.com/downloads)
3. Using a shell with Git (such as Git Bash), run the following command:

```
go get -u github.com/anaminus/rbxmk/rbxmk
```

If you installed Go correctly, this will install rbxmk to `$GOPATH/bin`,
which will allow you run it directly from a shell.

This document uses POSIX-style flags (`-f`, `--flag`), although windows-style
flags (`/f`, `/flag`) are possible when rbxmk is compiled for Windows. If you
are compiling for Windows, you may choose to force POSIX-style flags with the
`forceposix` build tag:

```
go get -u -tags forceposix github.com/anaminus/rbxmk/rbxmk
```

For more information, see the [go-flags](https://godoc.org/github.com/jessevdk/go-flags) package.

## Usage

See [USAGE.md](rbxmk/doc/USAGE.md) for an overview on how to use rbxmk. See
[DOCUMENTATION.md](rbxmk/doc/DOCUMENTATION.md) for full details on how rbxmk
works.
