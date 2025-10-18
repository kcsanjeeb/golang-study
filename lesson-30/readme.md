---
title: io.Reader and io.Writer Interfaces
description: Essential I/O interfaces for reading and writing data in Go
date: 2024-02-15
tags: ['go', 'io', 'reader', 'writer', 'interfaces']
---

# io.Reader and io.Writer Interfaces

## 📖 io.Reader Interface

### Definition
The `io.Reader` interface defines the contract for reading data from various sources.

### Interface Signature
~~~go
type Reader interface {
    Read(p []byte) (n int, err error)
}
~~~

### Key Behavior
- **Parameter**: Takes a byte slice to read data into
- **Return values**: 
  - `n`: Number of bytes successfully read
  - `err`: Error encountered during reading
- **Processing Rule**: Always process `n` bytes before handling the error
- **Sentinel Error**: `io.EOF` indicates end of data stream

### Common Implementations
- **`*os.File`**: Reading from files
- **`strings.Reader`**: Reading from strings
- **`bufio.Reader`**: Buffered reading for performance
- **`http.Request.Body`**: Reading HTTP request bodies
- **`bytes.Buffer`**: Reading from in-memory byte buffers

## ✍️ io.Writer Interface

### Definition
The `io.Writer` interface defines the contract for writing data to various destinations.

### Interface Signature
~~~go
type Writer interface {
    Write(p []byte) (n int, err error)
}
~~~

### Key Behavior
- **Parameter**: Takes source data as a byte slice
- **Return values**:
  - `n`: Number of bytes successfully written
  - `err`: Error encountered during writing
- **Error Handling**: Return error if written bytes < provided bytes
- **Immutability**: Provided byte slice should not be modified

### Common Implementations
- **`*os.File`**: Writing to files
- **`os.Stdout`**: Writing to standard output
- **`http.ResponseWriter`**: Writing HTTP response bodies
- **`bytes.Buffer`**: Writing to in-memory byte buffers
