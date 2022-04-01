gofmt -s -w .

d="/usr/local/go/src/srv"
mkdir -p "$d"
cp -r server "$d"
cp -r controllers "$d"
cp -r config "$d"
cp -r databases "$d"

go run main.go
