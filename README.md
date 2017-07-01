# rbxmk

`rbxmk` is a command-line tool for manipulating Roblox files.

The general workflow is that **inputs** are specified, transformed somehow,
then mapped to **outputs**.

[Lua](https://lua.org) scripts are used to perform actions. Scripts are run in
a stripped-down environment, with only a small set of functions available.

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

See [USAGE.md](USAGE.md) for details on how to use `rbxmk`.
