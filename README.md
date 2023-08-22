# flutter-webrtc-server-docker

[![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/r/iamriajul/flutter-webrtc-server) [![Docker Pulls](https://img.shields.io/docker/pulls/iamriajul/flutter-webrtc-server.svg?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/r/iamriajul/flutter-webrtc-server)  [![Docker Image Version](https://img.shields.io/docker/v/iamriajul/flutter-webrtc-server.svg?style=for-the-badge&logo=docker&logoColor=white&label=Docker%20Image%20Version)](https://hub.docker.com/r/iamriajul/flutter-webrtc-server) [![GitHub release (with filter)](https://img.shields.io/github/v/release/flutter-webrtc/flutter-webrtc-server?style=for-the-badge&logo=github&label=Flutter%20WebRTC%20Server)](https://github.com/flutter-webrtc/flutter-webrtc-server)

Docker Hub Tags: https://hub.docker.com/r/iamriajul/flutter-webrtc-server/tags

## Example docker-compose.yml

```yaml
version: '3.7'

services:
  agora-token:
    image: iamriajul/flutter-webrtc-server:latest # or specify a tag like iamriajul/flutter-webrtc-server:1.4.2
    ports:
      - 8080:8080 # the image exposes port 8080
    environment:
      - APP_ID=your-app-id
      - APP_CERTIFICATE=your-app-certificate
      - CORS_ALLOW_ORIGIN=your-allowed-origins
      # Note SERVER_PORT, PORT are not supported, as the image exposes port 8080
      # Passing these env variables will not work
    networks:
      - my-network
```
