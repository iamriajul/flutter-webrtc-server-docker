# flutter-webrtc-server-docker

[![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/r/iamriajul/flutter-webrtc-server) [![Docker Pulls](https://img.shields.io/docker/pulls/iamriajul/flutter-webrtc-server.svg?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/r/iamriajul/flutter-webrtc-server)  [![Docker Image Version](https://img.shields.io/docker/v/iamriajul/flutter-webrtc-server.svg?style=for-the-badge&logo=docker&logoColor=white&label=Docker%20Image%20Version)](https://hub.docker.com/r/iamriajul/flutter-webrtc-server) [![GitHub release (with filter)](https://img.shields.io/github/v/release/flutter-webrtc/flutter-webrtc-server?style=for-the-badge&logo=github&label=Flutter%20WebRTC%20Server)](https://github.com/flutter-webrtc/flutter-webrtc-server)

Docker Hub Tags: https://hub.docker.com/r/iamriajul/flutter-webrtc-server/tags

## Example docker-compose.yml

```yaml
version: '3.7'

services:
  agora-token:
    image: iamriajul/flutter-webrtc-server:latest # or specify a tag like iamriajul/flutter-webrtc-server:1.0
    ports:
      - 8086:8086 # the image exposes port 8086 for WebSocket Server.
      - 19302:19302 # the image exposes port 19302 for TURN/STUN Server.
    environment:
      # Optional Environment Variables
      - DOMAIN=demo.cloudwebrtc.com
      - CERT=configs/certs/cert.pem
      - KEY=configs/certs/key.pem
      - BIND=0.0.0.0
      - HTML_ROOT=web
      - PUBLIC_IP=127.0.0.1
      - REALM=flutter-webrtc
      # Note: PORT, TURN_PORT are fixed and cannot be changed by environment variables.
    networks:
      - my-network
```
