# Project Documentation

This folder is the guided study path for the distributed file system.

## Start Here

1. Read [Project Overview](overview.md) to understand the purpose and current behavior.
2. Read [System Architecture](architecture.md) to follow a file from a client call to disk and across the network.
3. Use the [File Descriptor](file-descriptor.md) as a map while opening source files.
4. Review [Important Takeaways](important-takeaways.md) for design decisions, risks, and next improvements.
5. Practice with [Interview Questions](interview-questions.md).

## Suggested Study Loop

For each feature, answer three questions:

- Which layer owns this behavior?
- What is the data format at this boundary?
- What happens when the peer, disk, or decoder fails?

Then run the smallest check that exercises it:

```powershell
go test ./...
mingw32-make build
```

## Current Scope

The repository is an educational distributed file-storage prototype. It currently demonstrates:

- Content-addressed file paths based on SHA-1.
- Local file reads and writes.
- AES-CTR streaming encryption and decryption.
- TCP peer connections with a small RPC envelope.
- File-store and file-fetch messages between nodes.
- Bootstrap connections between multiple demo servers.

It is not yet production-ready. Authentication, authenticated encryption, robust message framing, retries, replication policy, and concurrency protection need further work. Those gaps are intentionally listed in the architecture and takeaways documents.
