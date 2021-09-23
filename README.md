# atomicswap

This tool performs atomic filesystem swap operations on platforms that support
such functionality. At the moment, that includes:

- Linux (via `renameat2` with `RENAME_EXCHANGE`)
- macOS (via `renameatx_np` with `RENAME_SWAP`)

It does not (and will not) aim to support emulation of such functionality on
platforms that don't provide native support for atomic swap operations.

It's also worth noting that not all filesystems supported by these platforms
support atomic swap operations.


## Usage

```
$ atomicswap <first-path> <second-path>
```


## Requirements

Go 1.12 or later is required to build.


## Installation

The tool is a standard Go project and can be built/installed idiomatically, for
example:

```
$ git clone https://github.com/xenoscopic/atomicswap.git
$ cd atomicswap
$ go build .
```

or:

```
$ go install github.com/xenoscopic/atomicswap@latest
```
