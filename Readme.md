# Concurrent URL Checker with Logging and Reporting

## Overview

This is a simple program that takes a list of URLs, checks their status concurrently, logs the results, and generates a summary report.

## Running

```
go run main.go
```

## Testing

Run all tests

```
go test ./...
```

Run specific test

```
go test ./checker
```

## Purpose

- Demonstrates core Go features like concurrency(`goroutines`, `channels`) and error handling
- Shows practical usage of standard libraries (`net/http`, `flag`, `log`, etc...)
- Higlight clean code practices through modular design
- Includes testing which is essential for professional development
