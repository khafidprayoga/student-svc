FROM docker.io/golang:1.19 as builder

WORKDIR /student-service

ENV GO111MODULE=on GOBIN=/usr/local/bin
# Initialize dependency build
RUN go install github.com/bufbuild/buf/cmd/buf@v1.15.1 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

COPY . .
RUN buf generate proto/

RUN go mod tidy -v && go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo -o /out/student-service main.go

FROM docker.io/alpine:latest
COPY --from=builder /out/student-service start-service
CMD ["sh","-c","/start-service"]