---
title: Working with Time in Go
description: Time operations, measurements, and durations using the time package
date: 2024-02-17
tags: ['go', 'time', 'duration', 'measurement', 'timing']
---

# Working with Time in Go

## ⏰ Basic Time Operations

### Getting Current Time
```go
now := time.Now()
fmt.Println(now) // 2024-02-17 15:04:05.123456789 +0000 UTC
```

### Time Arithmetic
```go
// Add duration to time
twoHours := 2 * time.Hour
futureTime := now.Add(twoHours)

// Subtract duration from time  
pastTime := now.Add(-30 * time.Minute)

// Calculate time difference
duration := futureTime.Sub(now)
```

## ⏱️ Measuring Execution Time

### Using time.Since()
```go
start := time.Now()

// Code to measure
a := 1
for i := 0; i < 10000000; i++ {
    a++
}

elapsed := time.Since(start)
fmt.Println("Execution time:", elapsed)
```

### Using time.Now() Subtraction
```go
start := time.Now()

// Code block to measure

elapsed := time.Now().Sub(start)
```

## 📊 Common Time Durations

### Predefined Duration Constants
```go
time.Nanosecond
time.Microsecond
time.Millisecond
time.Second
time.Minute
time.Hour
```

### Creating Custom Durations
```go
duration := 2 * time.Hour
duration := 30 * time.Minute
duration := 500 * time.Millisecond
```