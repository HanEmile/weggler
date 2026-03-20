# weggler

A [weggli](https://github.com/weggli-rs/weggli) frontend

## Build + Run

```bash
go build -o weggler
./weggler -addr "127.0.0.1:8081" -max-jobs 2
```

## Delete all jobs:

```bash
$ rm data/jobs/*
```

## Reset statistics:

```bash
$ jq '(.[] | .avg_bytes_per_sec, .avg_lines_per_sec, .timing_samples) = 0' scripts.json
```

## Build for x86

```bash
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weggler
```
