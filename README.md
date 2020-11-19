# SimpleEcho

Using labstack/echo to echo the request.

## Cross Compiling on Windows 10 for Linux

```bash
set GOOS=linux
set GOARCH=amd64
go build
```

## Build Docker

```bash
docker build -t mikkoxu/simpleecho .
```

## Run with Docker

```bash
docker run -p 8080:8080 mikkoxu/simpleecho
```

```bash
docker run -d -p 8080:8080 mikkoxu/simpleecho
```
