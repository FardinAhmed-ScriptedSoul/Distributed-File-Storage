# Interview Questions

Use these questions to test whether you understand the project rather than memorizing names.

## System Understanding

1. What problem does this project solve?
   - It stores files locally and transfers them between cooperating nodes over TCP.

2. What is the responsibility of `FileServer`?
   - It coordinates the local `Store`, peer connections, application messages, and file streams.

3. Why is there a `Transport` interface?
   - To keep file-server logic independent from the concrete network implementation.

4. What is the difference between a message and a stream here?
   - A message is a small control command; a stream is a bounded sequence of raw file bytes.

5. Why is the file size sent before the file bytes?
   - The receiver needs a boundary so it can consume only this file and leave later protocol data untouched.

## Go and Networking

6. Why can one TCP `Read` return less data than requested?
   - TCP preserves order, not application message boundaries. A protocol must frame messages and read until the frame is complete.

7. What does `io.LimitReader` protect against in this project?
   - It limits a file transfer to the advertised byte count, preventing the receiver from consuming subsequent data.

8. Why must files and network connections be closed?
   - They hold operating-system resources and, on Windows, open file handles can prevent deletion.

9. What does embedding `net.Conn` in `TCPPeer` provide?
   - It lets a peer satisfy reader, writer, and address methods while adding project-specific `Send` and stream lifecycle behavior.

10. What does a handshake hook enable?
    - Future authentication or capability negotiation without changing transport setup code.

## Storage and Cryptography

11. Why transform a key into nested directories?
    - It distributes files across directories and avoids a single directory containing every object.

12. Is the current storage encrypted at rest?
    - No. The local `Store` writes plaintext. AES-CTR is used when file bytes are sent over the network.

13. Why is the IV sent with the ciphertext?
    - The decrypting side needs the same IV to reproduce the AES-CTR keystream. The IV is not secret, but it must be unique for encryption.

14. Why is AES-CTR alone insufficient for production?
    - It does not authenticate ciphertext, so tampering may go undetected. Use an AEAD construction.

15. Are MD5 and SHA-1 acceptable for security?
    - No. They are used here for deterministic paths and identifiers, not authentication or signatures.

## Debugging and Design

16. Why did the copied project fail to build initially?
    - Source files imported the upstream module path while `go.mod` declared the local module, and the test imported an undeclared assertion dependency.

17. Why should fixed sleeps not coordinate distributed requests?
    - Startup and network latency vary. Readiness signals, deadlines, request IDs, and response channels are more reliable.

18. Where can a data race occur?
    - `FileServer.peers` is modified when peers connect but read by broadcast and fetch paths without a consistent lock.

19. What happens if a peer sends a malformed size?
    - The current code trusts the decoded size. A production protocol must validate bounds, reject impossible values, and enforce deadlines.

20. How would you test a complete store/fetch round trip?
    - Start two transports on ephemeral ports, connect them, store a known byte sequence on one node, remove the local copy, fetch it from the other node, and compare bytes.

## Practical Exercise

Implement one improvement and explain its tradeoffs:

- Replace the fixed message buffer with a length-prefixed decoder.
- Add a request ID to `MessageGetFile` and return a response channel.
- Protect all peer-map snapshots with a read lock.
- Replace AES-CTR with AES-GCM.
- Add a graceful shutdown test.
