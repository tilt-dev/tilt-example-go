set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o build/tilt-example-go ./