FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Create a static binary for the server.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -a -installsuffix cgo -o /go/bin/server ./cmd/server/main.go

# Create a static binary for the configs.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -a -installsuffix cgo -o /go/bin/configs ./configs/main.go

# Create a minimal container to run a Golang static binary
FROM --platform=${TARGETPLATFORM:-linux/amd64} scratch

ENV GIN_MODE=release

# Copy our static executable.
COPY --from=builder /go/bin/server /server
COPY --from=builder /go/bin/configs /configs

EXPOSE 8086
EXPOSE 19302

ENTRYPOINT ["/configs && /server"]