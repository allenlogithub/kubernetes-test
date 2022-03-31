gofmt -s -w .

d="/usr/local/go/src/srv"
mkdir -p "$d"
cp -r server "$d"
cp -r controllers "$d"

CGO_ENABLED=0 go test -v srv/controllers
