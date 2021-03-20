pbspec
======

pbspec is a Go package to load the Protocol Buffer Language (proto files) for code generation.

```bash
go get h12.io/pbspec
```

A CLI tool `pbspec2json` is provided for exporting the ProtoBuf file descriptor set into JSON for inspection.

Usage:

```bash
pbspec2json -IPATH1 -IPATH2 ... PROTO_FILES
```
