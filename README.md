# weggler

A [weggli](https://github.com/weggli-rs/weggli) frontend

## Delete all jobs:

```bash
$ cat weggler.json | jq "del(.jobs)" > a.json && mv a.json weggler.json
```

## Reset statistics:

```bash
$ jq '(.[] | .avg_bytes_per_sec, .avg_lines_per_sec, .timing_samples) = 0' scripts.json
```

## Build for x86

```bash
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weggler
```
