gofmt -s -w .

d="/usr/local/go/src/srv"
mkdir -p "$d"
cp -r server "$d"
cp -r controllers "$d"

go run main.go
