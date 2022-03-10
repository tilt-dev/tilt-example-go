FROM golang:1.17

WORKDIR /src

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN go build -gcflags "all=-N -l" -o ./tilt-example-go ./

CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "./tilt-example-go"]
