$env:CGO_ENABLED=1; $env:CGO_CFLAGS="-IC:/tdlib/td/tdlib/include"; $env:CGO_LDFLAGS="-LC:/tdlib/td/tdlib/bin -ltdjson"; go build -trimpath -ldflags="-s -w" -o .\cmd\run.exe .\cmd\main.go
cd .\cmd
& .\run.exe