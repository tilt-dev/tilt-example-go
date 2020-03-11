FROM alpine
RUN apk add entr
WORKDIR /app
ADD web web
ADD build build
ENTRYPOINT find . | entr -r ./build/tilt-example-go
