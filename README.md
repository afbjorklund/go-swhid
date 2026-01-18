# Swhid

A Go library and CLI for generating and parsing SoftWare Hash IDentifiers (SWHIDs).

SWHIDs are persistent, intrinsic identifiers for software artifacts such as files, directories, commits, releases, and snapshots.
They are content-based identifiers that use Merkle DAGs for tamper-proof identification with built-in integrity verification.

This implementation follows the official [SWHID specification v1.2](https://www.swhid.org/specification) (ISO/IEC 18670:2025).
