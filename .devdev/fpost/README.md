# fpost
File sharing app

## Install
```bash
go install github.com/enuesaa/playground-2024/.devdev/fpost
```

## How to use
`fpost` trasnfers files in the current dir, from `source computer` to `destination computer`.

### On source computer
Run `fpost`

```console
cd /files-dir

# This serves on port 80
fpost
```

### On destination computer
```bash
curl 127.0.0.1 | sh
```
