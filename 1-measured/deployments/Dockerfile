FROM golang:1.14
WORKDIR /app
ADD . .
RUN go install ./
ENTRYPOINT tilt-example-go
