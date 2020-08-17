# Flutter Webapp server
### Features

- WebRTC server for flutter app
- Image to text translation

### Run Binary(Only Linux)

- Run

```bash
./bin/server-linux-amd64
```

Open https://0.0.0.0:8086.

### Compile from Source

- Install [tesseract-ocr](https://github.com/tesseract-ocr/tesseract/wiki)
- Go is required in the system
- Clone the repository, run `make`.
- Run `./bin/server-{platform}-{arch}`.
- Your server is live.

## Note

This example can only be used for LAN testing. If you need to use it in a production environment, you need more testing and and deploy an available turn server.
