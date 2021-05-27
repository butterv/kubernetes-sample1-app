FROM golang:1.16.4-alpine as builder

ENV GOPATH=/go \
    GO111MODULE=on \
    PROJECT_ROOT=/go/src/github.com/butterv/kubernetes-sample1-app

WORKDIR $PROJECT_ROOT

RUN wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.2/grpc_health_probe-linux-amd64 -O grpc_health_probe && \
    chmod +x grpc_health_probe

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-gateway -a ./app/cmd/gateway
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc-server -a ./app/cmd/server

# image for release
FROM gcr.io/distroless/base:latest
ENV BUILDER_ROOT /go/src/github.com/butterv/kubernetes-sample1-app
ENV PROJECT_ROOT /
COPY --from=builder $BUILDER_ROOT/grpc-gateway $PROJECT_ROOT/grpc-gateway
COPY --from=builder $BUILDER_ROOT/grpc-server $PROJECT_ROOT/grpc-server
COPY --from=builder $BUILDER_ROOT/grpc_health_probe /bin/grpc_health_probe
