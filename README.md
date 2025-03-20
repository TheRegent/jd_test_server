# jd_test_server

## go version

go 1.24

## Port

Application reads the port from `JD_TEST_PORT` environment variable.
If not set, it defaults to `8080`.

To run the application locally, run:

```bash
JD_TEST_PORT=8081 go run ./cmd/main.go
```

## Endpoints

There are 4 endpoints:

1. `GET /health` - returns 200 OK if the server is running.
2. `GET /crash` - crashes the server immediately.
3. `GET /load/cpu?duration=10s&intensity=5` - starts a CPU load for the specified duration and intensity.
4. `GET /load/memory?duration=10s&intensity=5` - starts a memory load for the specified duration and intensity.
+
