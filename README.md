# Orbit Benchmark Suite

This repository contains the benchmark and integration test suite used to validate **Orbit**.

> **Original Repository:** https://github.com/aayush4532/orbit

## Setup

1. Clone this repository.
2. Rename `.env.example` to `.env`.
3. Fill in the required environment variables.
4. Install dependencies.

```bash
go mod tidy
```

5. Start the Orbit server.

```bash
go run ./cmd/server
```

6. Open a new terminal and run the benchmark suite.

```bash
go run ./cmd/tests
```

The benchmark will execute the complete integration flow and generate the performance report automatically.
