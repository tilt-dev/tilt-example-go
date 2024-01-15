FROM golang:1.21
WORKDIR /app
ADD . .
RUN go install ./
ENTRYPOINT tilt-example-go
