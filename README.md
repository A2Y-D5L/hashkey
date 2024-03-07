# [A2Y-D5L](https://github.com/A2Y-D5L) / hashkey

`hashkey` provides a simple, concurrency-safe way to generate deterministic hash keys from strings. It uses the FNV-1a algorithm for hashing.

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
    "path/to/hashkey" // replace with the actual import path
)

func main() {
    hash, err := hashkey.String("example")
    if err != nil {
        fmt.Println("Error generating hash:", err)
        return
    }
    fmt.Println("Hexadecimal hash key:", hash)
}
```

#### Generate a uint64


```go
package main

import (
    "fmt"
    "github.com/A2Y-D5L/hashkey"
)

func main() {
    hash, err := hashkey.Uint64("example")
    if err != nil {
        fmt.Println("Error generating hash:", err)
        return
    }
    fmt.Println("Uint64 hash key:", hash)
}
```

## Concurrency-Safety

The `hashkey` library is designed to be safe for concurrent use. Each call to `String` or `Uint64` operates independently, using its instance of the FNV-1a hash function, ensuring no shared state between calls.
