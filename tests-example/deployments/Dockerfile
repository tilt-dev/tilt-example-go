FROM alpine
WORKDIR /app
ADD web web
ADD build build
ENTRYPOINT build/tilt-example-go
