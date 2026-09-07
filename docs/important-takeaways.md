# Important Takeaways

## Design Lessons

1. The system separates orchestration, storage, cryptography, and transport. That makes each layer replaceable, but the interfaces must describe the operations the upper layer actually needs.
2. `io.Reader`, `io.Writer`, and `io.Copy` make large-file transfers streaming-friendly instead of requiring the whole file in memory.
3. Content-addressed paths make a key map deterministically to a disk location and distribute files across directories.
4. The stream includes a file size so the receiver can use `io.LimitReader` and avoid consuming the next protocol message.
5. A buffered RPC channel decouples the TCP read loop from the file-server message loop.

## Security Lessons

- AES-CTR encrypts bytes but does not authenticate them. A production design should use an AEAD mode such as AES-GCM or ChaCha20-Poly1305 and verify integrity before accepting data.
- MD5 and SHA-1 are used as deterministic identifiers here, not as security boundaries. Do not use them for passwords, signatures, or trust decisions.
- The handshake is currently a no-op. Nodes are not authenticated, and any reachable client could attempt to join.
- Encryption keys are generated independently for each demo server. There is no key exchange or persistent key management.

## Reliability and Concurrency Risks

- `FileServer.peers` is written under `peerLock` in `OnPeer`, but several read paths iterate over it without the same lock. Concurrent map access can race or panic.
- `CloseStream` relies on a `WaitGroup` protocol. A malformed or unexpected stream marker can cause incorrect `Done` behavior.
- `Get` waits a fixed 500 milliseconds instead of correlating a response with a request ID.
- A broadcast fails on the first peer error and does not report partial success.
- The decoder assumes network reads align with logical messages. TCP is a byte stream, so robust framing needs explicit lengths and read-full loops.
- `DefaultDecoder` has a fixed 1028-byte message buffer and is not suitable for arbitrary-size payloads.
- The current demo storage root is derived from a listen address. Production code should use a platform-safe explicit directory configuration.

## Next Engineering Steps

1. Replace marker-only decoding with length-prefixed frames.
2. Add request IDs and response routing instead of sleeps.
3. Protect peer-map snapshots with a mutex or copy-on-write strategy.
4. Use AEAD encryption with nonce and authentication-tag handling.
5. Add peer authentication and authorization.
6. Add integration tests with temporary ports and temporary directories.
7. Define replication, duplicate-write, delete, and conflict semantics.
8. Add graceful shutdown that closes listeners, peers, and channels exactly once.
9. Replace demo `main.go` with a command-line configuration layer.
10. Add metrics and structured logs for transfer sizes, latency, and peer failures.

## Commands Worth Memorizing

```powershell
go test ./...
go test -race ./...
go vet ./...
gofmt -w .\*.go .\p2p\*.go
mingw32-make build
mingw32-make run
```
