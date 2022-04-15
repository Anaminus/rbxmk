<a id="user-content-rbxmk" href="#user-content-rbxmk">
	<img src="assets/logo-flat-name.png" alt="rbxmk logo"/>
</a>

**rbxmk** is a command-line tool for manipulating [Roblox][roblox] files.

rbxmk is useful for development workflows that involve the combination of many
separate files. If your project is organized as [Lua][lua] files for scripting
and model files for assets, rbxmk makes it simple to combine them into a final
product, be it a game, plugin, model, module, and so on. rbxmk is also suitable
for more simple actions, such as downloading models or publishing games.

[roblox]: https://corp.roblox.com
[lua]: https://lua.org

## Download
The current version of rbxmk is **<version>v0.9.1</version>**. The following
builds are available for download:

| Windows                     | Mac                | Linux                       |
|-----------------------------|--------------------|-----------------------------|
| **[Windows 64-bit][win64]** | **[macOS][macos]** | **[Linux 64-bit][linux64]** |
| **[Windows 32-bit][win32]** |                    | **[Linux 32-bit][linux32]** |

See the [Release page][release] for more information on the current version.

*rbxmk is fully featured, but thorough testing of all features is still a work
in progress. Please practice redundancy and use backups to reduce the risk of
data loss. Be sure to [report][issues] issues as you encounter them!*

[win64]: https://github.com/Anaminus/rbxmk/releases/download/v0.9.1/rbxmk-v0.9.1-windows-amd64.zip
[win32]: https://github.com/Anaminus/rbxmk/releases/download/v0.9.1/rbxmk-v0.9.1-windows-386.zip
[macos]: https://github.com/Anaminus/rbxmk/releases/download/v0.9.1/rbxmk-v0.9.1-darwin-amd64.zip
[linux64]: https://github.com/Anaminus/rbxmk/releases/download/v0.9.1/rbxmk-v0.9.1-linux-amd64.zip
[linux32]: https://github.com/Anaminus/rbxmk/releases/download/v0.9.1/rbxmk-v0.9.1-linux-386.zip
[source]: https://github.com/Anaminus/rbxmk/archive/v0.9.1.zip
[release]: https://github.com/Anaminus/rbxmk/releases/tag/v0.9.1
[issues]: https://github.com/Anaminus/rbxmk/issues

## Usage
rbxmk is a command-line tool, and so requires a [command-line interface][CLI] to
use.

rbxmk primarily uses [Lua][lua] scripts to produce and retrieve data, transform
it, and send it off to a variety of sources. The main subcommand is `run`, which
executes a script:

```bash
echo 'print("Hello world!")' > hello-world.lua
rbxmk run hello-world.lua
# Hello world!
```

The [Documentation page](doc/README.md) provides a complete reference on how
rbxmk is used, as well as the API of the Lua environment provided by rbxmk.

[CLI]: https://en.wikipedia.org/wiki/Command-line_interface

### Examples
The [examples](doc/examples) directory contains examples of rbxmk scripts.

- [Convert an asset URL in the clipboard to a model in the clipboard][copy-model]
- [Download an asset to a local file][download-asset]

[copy-model]: doc/examples/copy-model.rbxmk.lua
[download-asset]: doc/examples/download-asset.rbxmk.lua

## Installation
In addition to prebuilt releases, rbxmk can be installed manually.

1. [Install Go](https://golang.org/doc/install)
2. [Install Git](http://git-scm.com/downloads)
3. Using a shell with Git (such as Git Bash), run the following command:

```bash
go install github.com/anaminus/rbxmk/rbxmk@latest
```

If you installed Go correctly, this will install the latest version of rbxmk to
`$GOPATH/bin`, which will allow you run it directly from a shell.

A specific version of rbxmk may be installed by replacing `latest` with a
version number (e.g. `v0.9.1`).

### Development
To compile and install the bleeding-edge version, the best way is to clone the
repository:

```bash
git clone https://github.com/anaminus/rbxmk
cd rbxmk/rbxmk
go install
```

More information is available in the [INSTALL](INSTALL.md) document.

An effort is made to ensure that the latest commit will at least compile.
However, it is not guaranteed that everything will be in a production-ready
state.

## License
The source code for rbxmk is available under the [MIT license][mit].

[mit]: LICENSE
