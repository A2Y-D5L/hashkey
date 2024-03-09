# [A2Y-D5L](https://github.com/A2Y-D5L) / hashkey

[![Go Reference](https://pkg.go.dev/badge/github.com/A2Y-D5L/hashkey.svg)](https://pkg.go.dev/github.com/A2Y-D5L/hashkey)
[![Go Report Card](https://goreportcard.com/badge/github.com/A2Y-D5L/hashkey)](https://goreportcard.com/report/github.com/A2Y-D5L/hashkey)

`hashkey` provides a simple, concurrency-safe way to generate deterministic hash keys from strings. Hashes are produced using the FNV-1a algorithm implemented in the Go standard library at [`hash/fnv`](https://pkg.go.dev/hash/fnv).

##### NOTE:
>
> The FNV hash was designed for fast hash table and checksum use, not cryptography.
> 
>  — [Fowler–Noll–Vo hash function - Wikipedia](https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function#Non-cryptographic_hash)

## Installation

To use the `hashkey` library, first ensure you have Go installed on your machine. Then, import the library into your Go project:

```go
import "github.com/A2Y-D5L/hashkey"
```

## Usage

The library provides two functions:

- `String(s string) (string, error)`: Returns a hexadecimal string representation of the hash of the input string.
- `Uint64(s string) (uint64, error)`: Returns the raw `uint64` hash of the input string.

### Examples

#### Generate a string

```go
package main

import (
    "fmt"
    "github.com/A2Y-D5L/hashkey"
)

func main() {
    hash, err := hashkey.String("Hello there.")
    if err != nil {
        fmt.Println("Error generating hashkey:", err)
        return
    }
    fmt.Printf("String hash key: \"%s\"\n", hash) // -> String hash key: "892362dd056bbf55"
```

#### Generate a uint64

```go
package main

import (
    "fmt"
    "github.com/A2Y-D5L/hashkey"
)

func main() {
    hash, err := hashkey.Uint64("General Kenobi! You are a bold one.")
    if err != nil {
        fmt.Println("Error generating hashkey:", err)
        return
    }
    fmt.Printf("Uint64 hash key: %x\n", hash) // -> Uint64 hash key: 0x58d72f7088dae07
}
```

## Concurrency-Safety

The `hashkey` library is designed to be safe for concurrent use. Each call to `String` or `Uint64` operates independently, using its instance of the FNV-1a hash function, ensuring no shared state between calls.
