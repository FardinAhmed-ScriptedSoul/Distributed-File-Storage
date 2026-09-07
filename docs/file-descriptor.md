# File Descriptor

This is the source map for the repository. Read the files in this order when learning the system.

| File | Responsibility | What to study |
| --- | --- | --- |
| `main.go` | Demo entry point | Creates three nodes, starts them, stores and fetches sample files. |
| `server.go` | Distributed file-server orchestration | Node configuration, peer map, store/get messages, bootstrap, lifecycle. |
| `store.go` | Local filesystem storage | Path transformation, file creation, reads, writes, deletion, cleanup. |
| `crypto.go` | Streaming encryption helpers | Random IDs, key hashes, AES-CTR IV handling, reader-to-writer copying. |
| `p2p/transport.go` | Network abstraction | `Peer` and `Transport` contracts used by higher layers. |
| `p2p/tcp_transport.go` | TCP implementation | Listening, dialing, handshakes, peer registration, RPC read loop. |
| `p2p/message.go` | Transport event model | Marker bytes and the `RPC` envelope. |
| `p2p/encoding.go` | RPC decoder implementations | Gob decoding and marker-based stream detection. |
| `p2p/handshake.go` | Connection handshake hook | `HandshakeFunc` extension point and no-op default. |
| `store_test.go` | Storage tests | CAS path shape, write/read/delete behavior. |
| `crypto_test.go` | Crypto round-trip test | Encrypt then decrypt returns original bytes. |
| `p2p/tcp_transport_test.go` | TCP setup test | Transport configuration and listener startup. |
| `Makefile` | Developer commands | `build`, `run`, and `test` wrappers. |
| `.gitignore` | Generated-file rules | Ignores binaries, test output, editor metadata, and OS files. |
| `go.mod` | Module identity | Declares the local import path and Go version. |
| `README.md` | Short project entry point | Setup, commands, architecture links, and limitations. |

## Important Symbols

### `makeServer` in `main.go`
Builds a TCP transport and a `FileServer` with a storage root and path transform function.

### `FileServer.Store`
Writes a local copy, broadcasts metadata, then sends an encrypted stream to connected peers.

### `FileServer.Get`
Reads locally when possible. Otherwise broadcasts a request, receives a bounded encrypted stream, decrypts it, and reads the recovered local copy.

### `CASPathTransformFunc`
Turns a logical key into a deterministic directory path and filename.

### `copyEncrypt` and `copyDecrypt`
Define the stream format: a 16-byte IV followed by AES-CTR ciphertext/plaintext.

### `DefaultDecoder.Decode`
Consumes the first marker byte and decides whether the next data is an application message or a raw file stream.

### `TCPTransport.handleConn`
Runs the connection lifecycle: construct peer, handshake, register, decode RPCs, and publish them to the transport channel.

## Reading Exercise

Open `main.go` and trace one `s3.Store` call. Write down the next function before opening it. Continue through `server.go`, `store.go`, `crypto.go`, and `p2p/tcp_transport.go`. Then repeat with `s3.Get` after the local delete. This exercise exposes the real architecture faster than reading every helper in isolation.
