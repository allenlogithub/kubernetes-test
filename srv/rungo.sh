gofmt -s -w .

d="/usr/local/go/src/srv"
mkdir -p "$d"
cp -r server "$d"

go run main.go
