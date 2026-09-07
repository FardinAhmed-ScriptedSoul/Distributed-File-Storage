# Project Overview

## What the System Does

This project is a small distributed file system prototype written in Go. A file server stores file content on local disk and can exchange files with other file servers over TCP. Each node has:

- A local `Store` for persistent files.
- A `FileServer` that coordinates local storage and network requests.
- A `TCPTransport` that connects peers and delivers messages.
- A stream-encryption helper that encrypts file bytes while they cross the network.

The demo in `main.go` creates three nodes on ports `3000`, `7000`, and `5000`. The third node stores a file, removes its local copy, asks the network for it, decrypts the response, and prints the recovered content.

## Main Use Case

```text
client calls Store(key, data)
        |
        v
local Store writes encrypted? No: local Store writes plaintext
        |
        v
FileServer broadcasts metadata and streams encrypted bytes to peers
        |
        v
remote FileServer writes the received stream to its local Store

client calls Get(key)
        |
        +--> local copy exists: read local disk
        |
        +--> local copy missing: broadcast a request
                              |
                              v
                    a peer streams file size and encrypted bytes
                              |
                              v
                    local node decrypts and writes the file
```

Important: the current `Store` API writes plaintext to disk. Encryption is used for the network stream in `Store` and `Get` workflows, not as transparent at-rest encryption.

## How to Run

From PowerShell:

```powershell
gofmt -w .\*.go .\p2p\*.go
go test ./...
mingw32-make run
```

The demo is a long-running process. Stop it with `Ctrl+C`.

To check the TCP listener from another PowerShell terminal:

```powershell
Test-NetConnection localhost -Port 3000
```

A successful result has `TcpTestSucceeded : True`.

## Learning Vocabulary

- **Node:** one running `FileServer` process.
- **Peer:** one TCP connection to another node.
- **Transport:** the network implementation that accepts, dials, and decodes peers.
- **RPC:** one decoded transport event containing a sender, payload, and stream flag.
- **Message:** one gob-encoded application-level command such as store or get.
- **Stream:** raw file bytes that follow a stream marker and a file-size header.
- **Content-addressed path:** a deterministic disk path derived from a key hash.

## Current Reality

The project is useful for learning boundaries and data flow, but several production properties are not implemented yet:

- TCP messages do not have a robust length-delimited framing protocol.
- AES-CTR provides confidentiality but does not authenticate or detect tampering.
- The peer map is not consistently protected during all reads and writes.
- There is no durable metadata/index or conflict-resolution strategy.
- A request waits using a fixed sleep instead of a response correlation mechanism.
